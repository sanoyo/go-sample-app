package presentation

import (
	"github.com/labstack/echo"	
)

func Router() *echo.Echo {
	e := echo.New()
	return e
}