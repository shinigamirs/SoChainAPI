package infrastructure

import (
	"github.com/labstack/echo/v4"
	"nuri-challenge/sochainAPI/application"
	"nuri-challenge/sochainAPI/domain"
)

type SochainAPIController struct {
	sc *application.SochainClient
}

// NewSochainAPIController create a new SochainAPIController struct
func NewSochainAPIController(sc domain.APIClient) *SochainAPIController {
	return &SochainAPIController{application.NewSochainClient(sc)}
}

// GetTxDetailsController is the controller for get tx details
func (soAPIController *SochainAPIController) GetTxDetailsController(c echo.Context) error {
	return soAPIController.sc.GetTxDetailsHandler(c)
}

// GetBlockDetailsController is the controller for get blocks details
func (soAPIController *SochainAPIController) GetBlockDetailsController(c echo.Context) error {
	return soAPIController.sc.GetBlockDetailsHandler(c)
}
