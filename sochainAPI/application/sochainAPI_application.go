package application

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"nuri-challenge/sochainAPI/domain"
	"sync"
)

var supportedCrypto = map[string]bool{
	"BTC":  true,
	"LTC":  true,
	"DOGE": true,
}

type SochainClient struct {
	domain.APIClient
}

func NewSochainClient(sc domain.APIClient) *SochainClient {
	return &SochainClient{sc}
}

//GetBlockDetailsHandler get block details of given hash
func (sc *SochainClient) GetBlockDetailsHandler(c echo.Context) error {
	// get the param value and check if crypto is supported
	crypto := c.Param("crypto")
	blockHash := c.Param("blockHash")
	_, supported := supportedCrypto[crypto]
	if !supported {
		errMsg := fmt.Sprintf("%s is not supported", crypto)
		log.Error(errMsg)
		return echo.NewHTTPError(http.StatusMethodNotAllowed, errMsg)
	}

	// create and make request to get block details
	blockDetail, statusCode, err := sc.GetBlockDetails(blockHash, crypto)
	if err != nil {
		return echo.NewHTTPError(statusCode, err.Error())
	}

	// Creating Output JSON
	blockDetailOutput := domain.BlockDetailsOutput{
		NetworkCode:       blockDetail.Data.NetworkCode,
		BlockNumber:       blockDetail.Data.BlockNumber,
		Date:              blockDetail.Data.Date,
		PreviousBlockHash: blockDetail.Data.PreviousBlockHash,
		NextBlockHash:     blockDetail.Data.NextBlockHash,
		Size:              blockDetail.Data.Size,
	}

	// allocating 10 TxDetailsOutput to store txDetails
	blockDetailOutput.LastTenTxDetails = make([]domain.TxDetailsOutput, 10)
	var wg sync.WaitGroup
	for i := 0; i < len(blockDetail.Data.TxIDs) && i < 10; i++ {
		wg.Add(1)
		// To fetch the TxDetails concurrently
		go func(i int) {
			defer wg.Done()
			txDetails, _, _ := sc.GetTxDetails(blockDetail.Data.TxIDs[i], blockDetail.Data.NetworkCode)
			blockDetailOutput.LastTenTxDetails[i] = *txDetails
		}(i)
	}
	wg.Wait()
	return c.JSON(200, blockDetailOutput)
}

//GetTxDetailsHandler get transaction details of given transaction id
func (sc *SochainClient) GetTxDetailsHandler(c echo.Context) error {
	// get the param value and check if crypto is supported
	crypto := c.Param("crypto")
	txId := c.Param("transactionId")
	_, supported := supportedCrypto[crypto]
	if !supported {
		errMsg := fmt.Sprintf("%s is not supported", crypto)
		log.Error(errMsg)
		return echo.NewHTTPError(http.StatusMethodNotAllowed, errMsg)
	}

	txDetailsOutput, statusCode, err := sc.GetTxDetails(txId, crypto)
	if err != nil {
		return echo.NewHTTPError(statusCode, err.Error())
	}
	return c.JSON(200, txDetailsOutput)
}
