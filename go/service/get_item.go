package service

import (
	"mercari-build-training-2022/domain/item"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetItemResponse struct {
	Message string    `json:"message"`
	Item    item.Item `json:"items"`
}

func (s *projectService) GetItem(c echo.Context) error {
	// Get ID
	id := c.Param("itemId")

	item, err := s.itemRepo.GetItemById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GetItemResponse{Message: "id invalid"})
	}

	return c.JSON(http.StatusOK, GetItemResponse{
		Message: "hi",
		Item:    *item,
	})
}
