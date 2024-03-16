package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	apiV1Path = "/api/v1"
	authPath  = "/auth"
)

func New(handler *echo.Echo) {
	handler.GET(apiV1Path, info)
	auth := handler.Group(authPath)
	{
		newAuthRoutes(auth)
	}
}

// TODO: Заменить на Swagger
func info(c echo.Context) error {
	return c.HTML(http.StatusOK, `
		<!DOCTYPE html>
		<html>
		<head>
			<h1> Postl API V0.0.1 </h1>
    	</head>
		</html>
	`)
}
