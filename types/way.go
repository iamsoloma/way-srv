package types

import (
	"time"
)

type Block struct {
	ID       int       `json:"id"`
	Time_UTC time.Time `json:"time_utc"`
	PrevHash []byte    `json:"prev_hash"`
	Hash     []byte    `json:"hash"`
	Data     []byte    `json:"data"`
}
