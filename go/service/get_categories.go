package service

import (
	"mercari-build-training-2022/domain/category"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetCategoriesResponse struct {
	Message    string              `json:"message"`
	Categories []category.Category `json:"categories"`
}

func (s *projectService) GetCategories(c echo.Context) error {
	categories, err := s.catRepo.GetAllCategories(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, GetCategoriesResponse{Message: "Error fetching categories"})
	}

	return c.JSON(http.StatusOK, GetCategoriesResponse{
		Message:    "hi",
		Categories: categories,
	})
}
