package types

import "time"


type CreateBlockChainRequest struct {
	Chain string `json:"chain"`
	Genesis string `json:"genesis"`
	Time_UTC time.Time `json:"time_utc"`
}