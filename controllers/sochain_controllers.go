package controllers

import (
	"github.com/labstack/echo/v4"
	"nuri-challenge/sochainAPI/domain"
	"nuri-challenge/sochainAPI/infrastructure"
)

func AddSochainRoutes(e *echo.Echo) {
	t := e.Group("")
	gt := infrastructure.NewSochainClient(domain.BaseURL)
	t.GET("/currencies/:crypto/block-details/:blockHash", infrastructure.NewSochainAPIController(gt).GetBlockDetailsController)
	t.GET("/currencies/:crypto/tx-details/:transactionId", infrastructure.NewSochainAPIController(gt).GetTxDetailsController)
}
