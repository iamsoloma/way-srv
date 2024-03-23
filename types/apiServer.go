package types


type CreateBlockChainRequest struct {
	ChainName string `json:"chainName"`
	Genesis string `json:"genesis"`
}