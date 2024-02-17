package handlers

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"template-service-go/internal/app/instance"
)

type Handlers struct {
	instance *instance.Instance

	api *router.Group
}

func InitHandlers(r *router.Router, inst *instance.Instance) {
	h := &Handlers{
		instance: inst,
	}

	h.initApi(r)
}

// initApi Initialize api routes in router
func (h *Handlers) initApi(r *router.Router) {
	h.api = r.Group("/api")

	h.initApiApp()

	// Default handler if handler not found
	r.NotFound = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		whResponseLog(ctx)
	}
}

func (h *Handlers) handle(handler HandlerFunc) fasthttp.RequestHandler {
	return wrapHandler(handler, h.instance.Config)
}
