package types

import "time"

type CreateBlockChainRequest struct {
	ChainName string `json:"ChainName"`
	Genesis   string `json:"Genesis"`
}

type DeleteBlockChainRequest struct {
	ChainName string `json:"ChainName"`
}

type GetLastBlockRequest struct {
	ChainName string `json:"ChainName"`
}

type GetBlockByIDRequest struct {
	ChainName string `json:"ChainName"`
	ID        int    `json:"ID"`
}

type AddBlockRequest struct {
	ChainName string    `json:"ChainName"`
	Data      []byte    `json:"Data"`
	Time_UTC  time.Time `json:"Time_UTC"`
}
