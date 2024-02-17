package clickhouse

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
)

type (
	CH struct {
		Dsn string `json:"dsn"`
		DB  *ch.DB `json:"connect"`
		ctx context.Context

		poolTicker *time.Ticker

		queryPools []PoolModel
		mutex      sync.Mutex
	}

	PoolModel interface{}
)

func InitCH(dsn string, debug bool) (*CH, error) {
	var chC CH
	chC.Dsn = dsn
	chC.queryPools = make([]PoolModel, 0)
	chC.ctx = context.Background()

	db := ch.Connect(ch.WithDSN(dsn))
	chC.DB = db

	// DEBUG
	if debug {
		db.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))
		if err := db.Ping(chC.ctx); err != nil {
			panic(err)
		}
	}

	// Auto Migrate Models Tables
	chC.CreateTables(&Event{})

	chC.poolTicker = time.NewTicker(time.Second)

	return &chC, nil
}

func (chC *CH) CreateTables(models ...interface{}) {
	for _, model := range models {
		if _, err := chC.DB.NewCreateTable().Model(model).Exec(chC.ctx); err != nil {
		}
	}
}

func (chC *CH) Close() {
	chC.poolTicker.Stop()
	err := chC.DB.Close()
	if err != nil {
		return
	}
}

// AddQueryToPool Добавление запроса в соответствующий пул
func (chC *CH) AddQueryToPool(modelSrc PoolModel) {
	chC.mutex.Lock()
	chC.queryPools = append(chC.queryPools, modelSrc)
	chC.mutex.Unlock()
}

// FlushQueryPool Отправка запросов для каждой таблицы из пулов
func (chC *CH) FlushQueryPool() {
	for range chC.poolTicker.C {
		chC.mutex.Lock()

		for _, src := range chC.queryPools {
			if src != nil {
				chC.executeQuery(src)
			}
		}
		chC.queryPools = make([]PoolModel, 0)

		chC.mutex.Unlock()
	}
}

func (chC *CH) executeQuery(src PoolModel) {
	if _, err := chC.DB.NewInsert().Model(src).Exec(chC.ctx); err != nil {
		log.Printf("Error when inserting a model: %v\n", err)
	}
}
