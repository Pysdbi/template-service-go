package http

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"template-service-go/internal/app/instance"
	"template-service-go/internal/transport/http/handlers"
)

type HttpServer struct {
	Address string
	Router  *router.Router

	instance *instance.Instance
}

func InitHttpServer(instance *instance.Instance) *HttpServer {
	var hs HttpServer

	hs.Address = fmt.Sprintf("%s:%s", instance.Config.HTTP.Host, instance.Config.HTTP.Port)
	hs.Router = router.New()

	handlers.InitHandlers(hs.Router, instance)
	return &hs
}

func (hs HttpServer) Serve() error {
	err := fasthttp.ListenAndServe(hs.Address, hs.Router.Handler)
	if err != nil {
		return err
	}
	return nil
}
