package handlers

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"template-service-go/internal/app"
)

type Handlers struct {
	app *app.App

	api *router.Group
}

func InitHandlers(app *app.App, r *router.Router) {
	h := &Handlers{
		app: app,
	}

	h.initApi(r)
}

// initApi Initialize api routes in router
func (h *Handlers) initApi(r *router.Router) {
	h.api = r.Group("/api")

	h.initApiApp()
	h.initApiUser()

	// Default handler if handler not found
	r.NotFound = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		whResponseLog(ctx)
	}
}

func (h *Handlers) handle(handler HandlerFunc) fasthttp.RequestHandler {
	return wrapHandler(handler, h.app.Config)
}
