package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"nuri-challenge/sochainAPI/domain"
	"nuri-challenge/sochainAPI/infrastructure"
	"os"
	"path/filepath"
	"testing"
)

func TestSochainRoutes(t *testing.T) {
	e := echo.New()
	expectedFolder := filepath.Join("../resources", "testing")

	t.Run("TestSochainRoutesGetBlockDetailsReturnsJSON", func(t *testing.T) {

		//ARRANGE
		blockHash := "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf"
		crypto := "BTC"
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/currencies/:crypto/block-details/:blockHash")
		c.SetParamNames("crypto", "blockHash")
		c.SetParamValues(crypto, blockHash)
		soChainClient := infrastructure.NewSochainClient(domain.BaseURL)
		soChainAPIController := infrastructure.NewSochainAPIController(soChainClient)

		// Assertions
		expectedBody, err := os.ReadFile(filepath.Join(expectedFolder, "expectedGetBlockDetails.json"))
		assert.NoError(t, err)
		if assert.NoError(t, soChainAPIController.GetBlockDetailsController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, string(expectedBody), rec.Body.String())
		}
	})

	t.Run("TestSochainRoutesGetBlockDetailsWithInvalidBlockHashReturnsInvalidBlockhashError", func(t *testing.T) {

		//ARRANGE
		blockHash := "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2ff"
		crypto := "BTC"
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/currencies/:crypto/block-details/:blockHash")
		c.SetParamNames("crypto", "blockHash")
		c.SetParamValues(crypto, blockHash)
		soChainClient := infrastructure.NewSochainClient(domain.BaseURL)
		soChainAPIController := infrastructure.NewSochainAPIController(soChainClient)

		// Assertions
		expectedMessage := "valid block number or hash number is required"
		err := soChainAPIController.GetBlockDetailsController(c)
		if assert.Error(t, err) {
			resp, ok := err.(*echo.HTTPError)
			if ok {
				assert.Equal(t, http.StatusNotFound, resp.Code)
				assert.Equal(t, expectedMessage, resp.Message)
			}

		}

	})

	t.Run("TestSochainRoutesGetBlockDetailsWithUnsupportedCryptoReturnsErrorNotSupported", func(t *testing.T) {

		//ARRANGE
		blockHash := "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf"
		crypto := "GOLEM"
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/currencies/:crypto/block-details/:blockHash")
		c.SetParamNames("crypto", "blockHash")
		c.SetParamValues(crypto, blockHash)
		soChainClient := infrastructure.NewSochainClient(domain.BaseURL)
		soChainAPIController := infrastructure.NewSochainAPIController(soChainClient)

		// Assertions
		expectedMessage := fmt.Sprintf("%s is not supported", crypto)
		err := soChainAPIController.GetBlockDetailsController(c)
		if assert.Error(t, err) {
			resp, ok := err.(*echo.HTTPError)
			if ok {
				assert.Equal(t, http.StatusMethodNotAllowed, resp.Code)
				assert.Equal(t, expectedMessage, resp.Message)
			}

		}

	})

	t.Run("TestSochainRoutesGetTxDetailsReturnsJSON", func(t *testing.T) {

		//ARRANGE
		txId := "dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10"
		crypto := "BTC"
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/currencies/:crypto/tx-details/:transactionId")
		c.SetParamNames("crypto", "transactionId")
		c.SetParamValues(crypto, txId)
		soChainClient := infrastructure.NewSochainClient(domain.BaseURL)
		soChainAPIController := infrastructure.NewSochainAPIController(soChainClient)

		// Assertions
		expectedBody, err := os.ReadFile(filepath.Join(expectedFolder, "expectedGetTxIdDetails.json"))
		assert.NoError(t, err)
		if assert.NoError(t, soChainAPIController.GetTxDetailsController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, string(expectedBody), rec.Body.String())
		}
	})

	t.Run("TestSochainRoutesGetTxDetailsWithInvalidTxIdReturnsInvalidTxIdError", func(t *testing.T) {

		//ARRANGE
		txId := "dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a"
		crypto := "BTC"
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/currencies/:crypto/tx-details/:transactionId")
		c.SetParamNames("crypto", "transactionId")
		c.SetParamValues(crypto, txId)
		soChainClient := infrastructure.NewSochainClient(domain.BaseURL)
		soChainAPIController := infrastructure.NewSochainAPIController(soChainClient)

		// Assertions
		expectedMessage := "valid tx Id is required"
		err := soChainAPIController.GetTxDetailsController(c)
		if assert.Error(t, err) {
			resp, ok := err.(*echo.HTTPError)
			if ok {
				assert.Equal(t, http.StatusNotFound, resp.Code)
				assert.Equal(t, expectedMessage, resp.Message)
			}

		}

	})

	t.Run("TestSochainRoutesGetTxDetailsWithUnsupportedCryptoReturnsErrorNotSupported", func(t *testing.T) {

		//ARRANGE
		txId := "dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10"
		crypto := "GOLEM"
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/currencies/:crypto/tx-details/:transactionId")
		c.SetParamNames("crypto", "transactionId")
		c.SetParamValues(crypto, txId)
		soChainClient := infrastructure.NewSochainClient(domain.BaseURL)
		soChainAPIController := infrastructure.NewSochainAPIController(soChainClient)

		// Assertions
		expectedMessage := fmt.Sprintf("%s is not supported", crypto)
		err := soChainAPIController.GetTxDetailsController(c)
		if assert.Error(t, err) {
			resp, ok := err.(*echo.HTTPError)
			if ok {
				assert.Equal(t, http.StatusMethodNotAllowed, resp.Code)
				assert.Equal(t, expectedMessage, resp.Message)
			}

		}

	})
}
