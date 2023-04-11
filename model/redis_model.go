package model

import "time"

type SetDataRedis struct {
	Key  string
	Data interface{}
	Exp  time.Duration
}
