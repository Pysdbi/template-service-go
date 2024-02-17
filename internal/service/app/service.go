package service

import (
	"fmt"
	"template-service-go/internal/app/instance"
	def "template-service-go/internal/service"
)

var _ def.AppService = (*service)(nil)

type service struct {
	inst *instance.Instance
}

func NewService(inst *instance.Instance) *service {
	return &service{
		inst: inst,
	}
}

func (s *service) HealthCheck() (string, error) {
	fmt.Println(s.inst.Config.Name, s.inst.Config.Version)
	return "ok", nil
}

func (s *service) GetAppInfo() (string, error) {
	info := fmt.Sprintf("%s v%s", s.inst.Config.Name, s.inst.Config.Version)
	return info, nil
}
