package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"template-service-go/internal/config"
	"template-service-go/internal/pkg/http/response"
	"time"

	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

type HandlerFunc func(resp *response.Response)

// wrapHandler _
func wrapHandler(hf HandlerFunc, cnf *config.Config) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// Logging Request
		whRequestLog(ctx, cnf.App.Debug)

		resp := response.NewResponse(ctx)
		hf(resp)
		if !resp.CastStat {
			resp.Send()
		}

		// Logging Response
		if cnf.App.Debug {
			whResponseLog(ctx)
		}
	}
}

// =======================
// Debugging =============

func debug(in []byte) {
	var formattedBody interface{}
	if err := json.Unmarshal(in, &formattedBody); err == nil {
		if formatted, err := json.MarshalIndent(formattedBody, "", "    "); err == nil {
			fmt.Printf("Data:\n%s\n", string(formatted))
		} else {
			fmt.Printf("Debug data: %s\n", string(in))
		}
	} else {
		fmt.Printf("Debug data: %s\n", string(in))
	}
}

// Logging requests
func whRequestLog(ctx *fasthttp.RequestCtx, details bool) {
	var builder strings.Builder

	timeColor := color.New(color.FgWhite, color.Bold).SprintFunc()
	methodColor := color.New(color.FgBlue, color.Bold).SprintFunc()
	pathColor := color.New(color.FgMagenta, color.Bold).SprintFunc()

	builder.WriteString(fmt.Sprintf("%s %s %s\n",
		timeColor(time.Now().Format("2006-01-02 15:04:05")),
		methodColor(string(ctx.Method())),
		pathColor(string(ctx.Path())),
	))

	if details {
		requestBody := ctx.PostBody()
		if len(requestBody) > 0 {
			var formattedBody interface{}
			if err := json.Unmarshal(requestBody, &formattedBody); err == nil {
				if formatted, err := json.MarshalIndent(formattedBody, "", "    "); err == nil {
					builder.WriteString(fmt.Sprintf("Body:\n%s\n", string(formatted)))
				} else {
					builder.WriteString(fmt.Sprintf("Body: %s\n", string(requestBody)))
				}
			} else {
				builder.WriteString(fmt.Sprintf("Body: %s\n", string(requestBody)))
			}
		}
	}

	fmt.Print(builder.String())
}

// Logging responses
func whResponseLog(ctx *fasthttp.RequestCtx) {
	var builder strings.Builder

	// Определение функций для цветного вывода
	timeColor := color.New(color.FgWhite, color.Bold).SprintFunc()
	methodColor := color.New(color.FgBlue, color.Bold).SprintFunc()
	pathColor := color.New(color.FgMagenta, color.Bold).SprintFunc()
	statusColorFunc := color.New(color.FgCyan, color.Bold).SprintFunc()

	statusCode := ctx.Response.StatusCode()
	statusText := fasthttp.StatusMessage(statusCode)

	switch {
	case statusCode >= 200 && statusCode < 300:
		statusColorFunc = color.New(color.FgGreen, color.Bold).SprintFunc()
	case statusCode >= 300 && statusCode < 400:
		statusColorFunc = color.New(color.FgYellow, color.Bold).SprintFunc()
	case statusCode >= 400:
		statusColorFunc = color.New(color.FgRed, color.Bold).SprintFunc()
	}

	// Форматирование и запись строки
	builder.WriteString(fmt.Sprintf("%s %s %s  code=%s msg=%s\n",
		timeColor(time.Now().Format("2006-01-02 15:04:05")),
		methodColor(string(ctx.Method())),
		pathColor(string(ctx.Path())),
		statusColorFunc(statusCode),
		statusColorFunc(statusText),
	))

	fmt.Print(builder.String())
}
