package audit

import "time"

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}
