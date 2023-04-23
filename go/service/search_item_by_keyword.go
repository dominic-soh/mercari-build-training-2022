package service

import (
	"mercari-build-training-2022/domain/item"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SearchItemResponse struct {
	Message string      `json:"message"`
	Items   []item.Item `json:"items"`
}

func (s *projectService) SearchItemByKeyword(c echo.Context) error {
	// Get form values
	keyword := c.QueryParam("keyword")

	items, err := s.itemRepo.GetItemsByName(c.Request().Context(), keyword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, SearchItemResponse{Message: "Issue with finding items"})
	}

	return c.JSON(http.StatusOK, items)
}
