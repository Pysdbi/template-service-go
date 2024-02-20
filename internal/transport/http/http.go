package http

import (
	"fmt"
	"template-service-go/internal/app"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"template-service-go/internal/transport/http/handlers"
)

type HttpServer struct {
	Address string
	Router  *router.Router

	app *app.App
}

func InitHttpServer(app *app.App) *HttpServer {
	var hs HttpServer

	hs.Address = fmt.Sprintf("%s:%s", app.Config.HTTP.Host, app.Config.HTTP.Port)
	hs.Router = router.New()

	handlers.InitHandlers(app, hs.Router)
	return &hs
}

func (hs HttpServer) Serve() error {
	err := fasthttp.ListenAndServe(hs.Address, hs.Router.Handler)
	if err != nil {
		return err
	}
	return nil
}
