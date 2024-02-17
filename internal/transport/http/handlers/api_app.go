package handlers

import (
	"fmt"
	"template-service-go/internal/pkg/http/response"
)

func (h *Handlers) initApiApp() {
	h.api.GET("/healthcheck", h.handle(func(resp *response.Response) {
		appInfo := fmt.Sprintf("%s v%s", h.instance.Config.Name, h.instance.Config.Version)
		resp.SetValue(appInfo)
	}))
}
