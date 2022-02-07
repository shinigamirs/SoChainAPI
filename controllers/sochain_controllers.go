package controllers

import (
	"github.com/labstack/echo/v4"
	"nuri-challenge/sochainAPI/domain"
	"nuri-challenge/sochainAPI/infrastructure"
)

func AddSochainRoutes(e *echo.Echo) {
	t := e.Group("")
	sc := infrastructure.NewSochainClient(domain.BaseURL)
	t.GET("/currencies/:crypto/block-details/:blockHash", infrastructure.NewSochainAPIController(sc).GetBlockDetailsController)
	t.GET("/currencies/:crypto/tx-details/:transactionId", infrastructure.NewSochainAPIController(sc).GetTxDetailsController)
}
