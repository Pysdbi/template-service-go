package clickhouse

import (
	"time"
)

type Event struct {
	EventTime  time.Time `ch:"EventTime"`                                                                   // Время события
	Event      string    `ch:"Event"`                                                                       // Тип события
	Message    *string   `ch:"Message"`                                                                     // Сообщение о событии
	Level      string    `ch:"Level,type:Enum8('DEBUG'=1, 'INFO'=2, 'WARNING'=3, 'ERROR'=4, 'CRITICAL'=5)"` // Уровень важности события
	EntityID   *string   `ch:"EntityID"`                                                                    // Идентификатор сущности
	EntityName *string   `ch:"EntityName"`                                                                  // Название сущности
}

func (e *Event) Write(ch *CH) {
	ch.AddQueryToPool(e)
}
