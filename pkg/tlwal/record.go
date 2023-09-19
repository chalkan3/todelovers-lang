package wal

import "time"

type Data struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Record struct {
	Timestamp time.Time `json:"timestamp"`
	Table     string    `json:"table"`
	Operation string    `json:"operation"`
	Data      *Data     `json:"data"`
}

func (r *Record) Key() string {
	return r.Data.Key
}
