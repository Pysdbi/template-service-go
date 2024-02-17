package instance

import (
	"template-service-go/internal/config"
	"template-service-go/internal/domain/clickhouse"
	"template-service-go/internal/domain/minio"
	"template-service-go/internal/domain/pgsql"
	"template-service-go/internal/transport/amqp"
)

type Instance struct {
	Config     *config.Config `json:"config"`
	Database   *pgsql.DB      `json:"database"`
	Clickhouse *clickhouse.CH `json:"clickhouse"`
	Amqp       *amqp.Amqp     `json:"amqp"`
	Minio      *minio.Minio   `json:"minio"`
}
