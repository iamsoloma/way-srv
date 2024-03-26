package types

type CreateBlockChainRequest struct {
	ChainName string `json:"chainName"`
	Genesis   string `json:"genesis"`
}

type DeleteBlockChainRequest struct {
	ChainName string `json:"chainName"`
}

type GetLastBlockRequest struct {
	ChainName string `json:"chainName"`
}

type GetBlockByIDRequest struct {
	ChainName string `json:"chainName"`
	ID        int    `json:"id"`
}

type AddBlockRequest struct {
	ChainName string `json:"chainName"`
	Data      []byte `json:"data"`
}
