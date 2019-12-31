package presentation

import (
	"github.com/labstack/echo"

	"github.com/go-sample-app/app/presentation/handler"
)

func Router() *echo.Echo {
	e := echo.New()
	e.GET("/contents/:contentID", handler.ShowContent)

	return e
}
