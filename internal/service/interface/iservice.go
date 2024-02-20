package iservice

type Services interface {
	GetAppService() AppService
	GetUserService() UserService
}

type AppService interface {
	Healthcheck() (bool, error)
	GetAppInfo() (string, error)
}

type UserService interface {
	CreateUser() int
}
