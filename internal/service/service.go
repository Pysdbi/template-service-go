package service

import (
	"template-service-go/internal/app"

	iservice "template-service-go/internal/service/interface"

	appService "template-service-go/internal/service/app"
	userService "template-service-go/internal/service/user"
)

func InitServices(app *app.App) *Services {
	services := &Services{
		AppService:  appService.NewService(app),
		UserService: userService.NewService(app),
	}
	return services
}

var _ iservice.Services = (*Services)(nil)

type Services struct {
	AppService  iservice.AppService
	UserService iservice.UserService
}

func (s Services) GetAppService() iservice.AppService {
	return s.AppService
}

func (s Services) GetUserService() iservice.UserService {
	return s.UserService
}
