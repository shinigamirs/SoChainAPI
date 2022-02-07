package domain

// GetBlockDetailResp struct to store response of GetBlockDetails
type GetBlockDetailResp struct {
	Status string `json:"status"`
	Data   struct {
		NetworkCode       string   `json:"network"`
		BlockNumber       int64    `json:"block_no"`
		Date              Time     `json:"time"`
		PreviousBlockHash string   `json:"previous_blockhash"`
		NextBlockHash     string   `json:"next_blockhash"`
		Size              int64    `json:"size"`
		TxIDs             []string `json:"txs"`
	} `json:"data"`
}

// GetTxDetailResp struct to store response of GetTxDetails
type GetTxDetailResp struct {
	Status string          `json:"status"`
	Data   TxDetailsOutput `json:"data"`
}

//BlockDetailsOutput output struct for GetBlockDetails endpoint
type BlockDetailsOutput struct {
	NetworkCode       string            `json:"network"`
	BlockNumber       int64             `json:"block_no"`
	Date              Time              `json:"time"`
	PreviousBlockHash string            `json:"previous_blockhash"`
	NextBlockHash     string            `json:"next_blockhash"`
	Size              int64             `json:"size"`
	LastTenTxDetails  []TxDetailsOutput `json:"LastTenTxDetails"`
}

//TxDetailsOutput output struct for GetTxDetails endpoint
type TxDetailsOutput struct {
	TxID      string `json:"txId"`
	Date      Time   `json:"time"`
	Fee       string `json:"fee"`
	SentValue string `json:"sent_value"`
}

const (
	BaseURL = "https://sochain.com/api/v2"
)
