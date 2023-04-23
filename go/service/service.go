package service

import (
	"mercari-build-training-2022/infrastructure/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type projectService struct {
	db       *gorm.DB
	itemRepo repository.ItemsRepository
	catRepo  repository.CategoryRepository
}

type ProjectService interface {
	CreateItemDB(c echo.Context) error
	GetAllItems(c echo.Context) error
	GetCategories(c echo.Context) error
	GetItem(c echo.Context) error
	SearchItemByKeyword(c echo.Context) error
	Root(c echo.Context) error
	GetImage(c echo.Context) error
}

func NewService(db *gorm.DB, itemsRepo repository.ItemsRepository, catRepo repository.CategoryRepository) ProjectService {
	return &projectService{
		db:       db,
		itemRepo: itemsRepo,
		catRepo:  catRepo,
	}
}
