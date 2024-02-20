package handlers

import (
	"template-service-go/internal/pkg/http/response"
)

func (h *Handlers) initApiApp() {
	appService := h.app.Services.GetAppService()

	h.api.GET("/healthcheck", h.handle(func(resp *response.Response) {
		ok, err := appService.Healthcheck()
		if err != nil {
			resp.SetError(5, err.Error())
		}

		resp.SetValue(ok)
	}))

	h.api.GET("/app-info", h.handle(func(resp *response.Response) {
		info, err := appService.GetAppInfo()
		if err != nil {
			resp.SetError(5, err.Error())
		}

		resp.SetValue(info)
	}))
}
