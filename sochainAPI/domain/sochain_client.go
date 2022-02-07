package domain

type APIClient interface {
	GetTxDetails(txId string, crypto string) (*TxDetailsOutput, int, error)
	GetBlockDetails(blockhash string, crypto string) (*GetBlockDetailResp, int, error)
}
