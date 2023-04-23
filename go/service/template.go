package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloWorldResponse struct {
	Message string `json:"message"`
}

func (s *projectService) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, HelloWorldResponse{Message: "Hello, world!"})
}
