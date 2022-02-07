package controllers

import (
	"github.com/labstack/echo/v4"
)

//InitEcho returns an echo struct with validator and session middleware
func InitEcho() *echo.Echo {
	e := echo.New()
	return e
}

//AddRoutes add the endpoint
func AddRoutes(e *echo.Echo) {
	AddSochainRoutes(e)
}
