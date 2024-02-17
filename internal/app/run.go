package app

import (
	"github.com/fatih/color"
	"template-service-go/internal/config"
	"template-service-go/internal/domain/clickhouse"
	"template-service-go/internal/domain/minio"
	"template-service-go/internal/domain/pgsql"
	"template-service-go/internal/transport/amqp"
	"template-service-go/internal/transport/http"
)

func Run() {

	// create app =================
	app := NewApp()

	// init Config ================
	conf, err := config.NewConfig()
	if err != nil {
		app.LogErr("error init config", err)
	}
	app.Config = conf

	//app.Debug()

	// init DB ============
	pg, err := pgsql.InitDB(app.Config.Dsn.Database)
	if err != nil {
		app.LogErr("error init gorm", err)
	}
	app.Database = pg

	// init Clickhouse =========
	ch, err := clickhouse.InitCH(app.Config.Dsn.Clickhouse, app.Config.App.Debug)
	if err != nil {
		app.LogErr("error init clickhouse", err)
	}
	app.Clickhouse = ch
	go ch.FlushQueryPool()

	// init S3/Minio ============
	s3, err := minio.InitMinio(app.Config.Minio.Host, app.Config.Minio.AccessKey, app.Config.Minio.SecretKey, false)
	if err != nil {
		app.LogErr("error init s3 minio", err)
	}
	app.Minio = s3

	// init Broker ================
	am, err := amqp.InitAMQP(app.Config.Dsn.Amqp)
	if err != nil {
		app.LogErr("error init amqp", err)
	}
	app.Amqp = am

	// init Services ============
	// TODO: appServices := service.InitServices(app.GetInstance())

	// init http server ============
	httpServer := http.InitHttpServer(app.GetInstance())
	go func() {
		err := httpServer.Serve()
		if err != nil {
			app.LogErr("error with httpServer.ListenAndServe: ", err)
		}
	}()

	// Show info about service
	app.PrintAppInfo([]InfoBlock{
		{
			title: "HTTP server",
			params: map[string]interface{}{
				"address": color.New(color.Underline).Sprintf("%s:%s", app.Config.HTTP.Host, app.Config.HTTP.Port),
				"status":  color.New(color.FgGreen).Sprint("ok"),
			},
		},
	})

	// forever wait ===============
	<-app.Close
	app.GracefulShutdown()

}
