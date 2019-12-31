package handler

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/go-sample-app/app/application/usecase"
)

func ShowContent(c echo.Context) error {
	contentID := c.Param("contentID")

	content, err := usecase.ShowContent(contentID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"string": "failed", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, content)
}
