package user

import (
	"math/rand"
	"template-service-go/internal/app"
	iservice "template-service-go/internal/service/interface"
)

func NewService(app *app.App) *Service {
	return &Service{
		app: app,
	}
}

var _ iservice.UserService = (*Service)(nil)

type Service struct {
	app *app.App
}

func (s Service) CreateUser() int {
	return rand.Int()
}
