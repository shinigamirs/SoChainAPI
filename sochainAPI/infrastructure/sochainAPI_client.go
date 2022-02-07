package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"nuri-challenge/sochainAPI/domain"
)

type SochainAPIClient struct {
	url string
}

func NewSochainClient(url string) SochainAPIClient {
	return SochainAPIClient{url: url}
}

//GetTxDetails call http request to sochainapi to fetch tx details and return TxDetailsOutput, statusCode, error
func (sc SochainAPIClient) GetTxDetails(txId string, crypto string) (*domain.TxDetailsOutput, int, error) {

	//creating url with crypto and txId
	url := fmt.Sprintf("%s/tx/%s/%s", sc.url, crypto, txId)
	res, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return nil, res.StatusCode, err
	}
	if res.StatusCode != http.StatusOK {
		// This to handle if transaction id is not valid sochain returns with 500 status code
		if res.StatusCode == http.StatusNotFound || res.StatusCode == http.StatusInternalServerError {
			return nil, http.StatusNotFound, errors.New("valid tx Id is required")
		} else {
			return nil, res.StatusCode, errors.New("something went wrong")
		}
	}

	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	//storing response in txDetails
	var txDetails domain.GetTxDetailResp
	err = json.Unmarshal(respBody, &txDetails)
	if err != nil {
		log.Panic(err)
	}
	return &txDetails.Data, res.StatusCode, nil
}

// GetBlockDetails calls request to sochainapi to fetch block details
func (sc SochainAPIClient) GetBlockDetails(blockHash string, crypto string) (*domain.GetBlockDetailResp, int, error) {

	// creating url with crypto and blockHash
	url := fmt.Sprintf("%s/get_block/%s/%s", sc.url, crypto, blockHash)

	res, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return nil, res.StatusCode, err
	}
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, res.StatusCode, errors.New("valid block number or hash number is required")
		} else {
			return nil, res.StatusCode, errors.New("something went wrong")
		}
	}
	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
	}

	// storing response in blockDetail
	var blockDetail domain.GetBlockDetailResp
	err = json.Unmarshal(respBody, &blockDetail)
	if err != nil {
		log.Panic(err)
	}
	return &blockDetail, res.StatusCode, nil
}
