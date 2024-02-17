package response

import (
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/valyala/fasthttp"
)

type ErrorItem struct {
	Key     int    `json:"key"`
	Message string `json:"message"`
}

type Response struct {
	Ctx         *fasthttp.RequestCtx `json:"-"`
	CastStat    bool                 `json:"-"`
	Status      bool                 `json:"status"`
	Errors      []ErrorItem          `json:"errors"`
	Values      []interface{}        `json:"values,omitempty" `
	Value       interface{}          `json:"value,omitempty"`
	Code        int                  `json:"-"`
	TmRequest   string               `json:"tm_req"`
	TmRequestSt time.Time            `json:"-"`

	isValueArray bool
}

func NewResponse(reqCtx *fasthttp.RequestCtx) *Response {
	return &Response{
		TmRequestSt:  time.Now(),
		Code:         200,
		Ctx:          reqCtx,
		isValueArray: true,
	}
}

func (r *Response) SetError(key int, mess string) *Response {
	r.Errors = append(r.Errors, ErrorItem{
		Key:     key,
		Message: mess,
	})
	return r
}

func (r *Response) SetCode(in int) *Response {
	r.Code = in
	return r
}

func (r *Response) SetValue(val interface{}) *Response {
	r.Value = val
	r.isValueArray = false
	return r
}

func (r *Response) SetValues(vals interface{}) *Response {
	r.Values = r.InterfaceSlice(vals)
	return r
}

func (r *Response) CastJson(in interface{}) {
	bts, err := json.Marshal(in)
	if err != nil {
		log.Println("error set cast json")
		return
	}
	r.Ctx.Request.Header.SetContentType("application/json")
	r.CastStat = true
	r.Ctx.Write(bts)
}

func (r *Response) SetAcceptJobHeaders() {
	r.Ctx.SetContentType("application/json")
	r.Ctx.SetStatusCode(r.Code)
}

func (r *Response) CastJsonClear(in []byte) {
	r.SetAcceptJobHeaders()
	r.CastStat = true
	r.Ctx.Response.AppendBody(in)
}

func (r *Response) NoContent() {
	r.SetAcceptJobHeaders()
	r.CastStat = true
}

func (r *Response) InterfaceSlice(in interface{}) []interface{} {
	s := reflect.ValueOf(in)
	if s.Kind() != reflect.Slice {
		return []interface{}{}
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

func (r *Response) FormResponse() *Response {
	if len(r.Errors) > 0 {
		r.Status = false
	} else {
		r.Status = true
	}

	// not null array json
	if len(r.Values) == 0 && r.isValueArray {
		r.Values = []interface{}{}
	}

	// not null array json
	if len(r.Errors) == 0 {
		r.Errors = []ErrorItem{}
	}
	return r
}

func (r *Response) Send() {
	r.Ctx.Request.Header.SetContentType("application/json")
	r.Ctx.Write(r.Json())
}

func (r *Response) Json() []byte {
	r.TmRequest = time.Now().Sub(r.TmRequestSt).String()
	if bts, err := json.Marshal(r.FormResponse()); err == nil {
		return bts
	}
	return []byte{}
}
