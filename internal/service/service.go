package service

import (
	"template-service-go/internal/app/instance"
	appService "template-service-go/internal/service/app"
)

type Services struct {
	AppService AppService
}

type AppService interface {
	HealthCheck() (string, error)
	GetAppInfo() (string, error)
}

func InitServices(inst *instance.Instance) *Services {
	services := &Services{
		AppService: appService.NewService(inst),
	}
	return services
}
