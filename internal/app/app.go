package app

import (
	"log"
	"template-service-go/internal/app/instance"
	"template-service-go/internal/config"
	"template-service-go/internal/domain/clickhouse"
	"template-service-go/internal/domain/minio"
	"template-service-go/internal/domain/pgsql"
	"template-service-go/internal/transport/amqp"
)

type App struct {
	Config *config.Config `json:"config"`

	// Databases
	Database   *pgsql.DB      `json:"database"`
	Clickhouse *clickhouse.CH `json:"clickhouse"`

	Amqp  *amqp.Amqp   `json:"amqp"`
	Minio *minio.Minio `json:"minio"`

	Close chan bool `json:"close"`
}

func NewApp() *App {
	app := &App{Close: make(chan bool)}
	return app
}

func (a *App) LogErr(errs ...interface{}) {
	log.Fatal(errs)
}

func (a *App) GetInstance() *instance.Instance {
	return &instance.Instance{
		Config:     a.Config,
		Database:   a.Database,
		Clickhouse: a.Clickhouse,
		Amqp:       a.Amqp,
		Minio:      a.Minio,
	}
}

// GracefulShutdown close all
func (a *App) GracefulShutdown() {
	log.Println("Graceful Shutdown App")
}
