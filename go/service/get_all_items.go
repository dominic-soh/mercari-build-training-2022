package service

import (
	"mercari-build-training-2022/domain/item"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetAllItemsResponse struct {
	Message string      `json:"message"`
	Items   []item.Item `json:"items"`
}

func (s *projectService) GetAllItems(c echo.Context) error {
	items, err := s.itemRepo.GetAllItems(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, AddItemResponse{Message: "No image uploaded"})
	}

	return c.JSON(http.StatusOK, GetAllItemsResponse{
		Message: "hi",
		Items:   items,
	})
}
