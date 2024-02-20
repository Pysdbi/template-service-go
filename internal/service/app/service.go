package app

import (
	"fmt"

	iservice "template-service-go/internal/service/interface"

	"template-service-go/internal/app"
)

func NewService(app *app.App) *Service {
	return &Service{
		app: app,
	}
}

type Service struct {
	app *app.App
}

var _ iservice.AppService = (*Service)(nil)

func (s *Service) Healthcheck() (bool, error) {
	return true, nil
}

func (s *Service) GetAppInfo() (string, error) {
	info := fmt.Sprintf(
		"%s v%s",
		s.app.Config.Name,
		s.app.Config.Version,
	)
	return info, nil
}
