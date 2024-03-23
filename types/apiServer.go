package types

type CreateBlockChainRequest struct {
	ChainName string `json:"chainName"`
	Genesis   string `json:"genesis"`
}

type GetLastBlock struct {
	ChainName string `json:"chainName"`
}

type GetBlockByID struct {
	ChainName string `json:"chainName"`
	ID        int    `json:"id"`
}

type AddBlock struct {
	ChainName string `json:"chainName"`
	Data      []byte `json:"data"`
}
