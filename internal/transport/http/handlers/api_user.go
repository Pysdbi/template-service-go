package handlers

import (
	"template-service-go/internal/pkg/http/response"
)

func (h *Handlers) initApiUser() {
	userService := h.app.Services.GetUserService()

	h.api.GET("/user/create", h.handle(func(resp *response.Response) {
		userId := userService.CreateUser()

		resp.SetValue(userId)
	}))
}
