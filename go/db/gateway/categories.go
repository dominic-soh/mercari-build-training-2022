package gateway

import (
	"context"
	"errors"
	"mercari-build-training-2022/db/models"
	"mercari-build-training-2022/domain/category"

	"gorm.io/gorm"
)

type CategoriesRepository struct {
	client *gorm.DB
}

func NewCategoriesRepository(client *gorm.DB) *CategoriesRepository {
	return &CategoriesRepository{
		client: client,
	}
}

func (cr *CategoriesRepository) GetAllCategories(ctx context.Context) ([]category.Category, error) {
	var categories []category.Category
	results := cr.client.Find(&categories)
	if results.Error != nil {
		return nil, results.Error
	}

	return categories, nil
}

func (cr *CategoriesRepository) CreateCategory(ctx context.Context, name string) (*category.Category, error) {
	newCategory := &models.Category{
		Name: name,
	}
	results := cr.client.Create(newCategory)
	if results.Error != nil {
		return nil, results.Error
	}

	return &category.Category{
		ID:   newCategory.ID,
		Name: newCategory.Name,
	}, nil
}

func (cr *CategoriesRepository) GetCategory(ctx context.Context, name string) (*category.Category, error) {
	var cat models.Category
	results := cr.client.Find(&cat).Where("name = ?", name)
	if results.Error != nil && !errors.Is(results.Error, gorm.ErrRecordNotFound) {
		return nil, results.Error
	}

	return &category.Category{
		ID:   cat.ID,
		Name: cat.Name,
	}, nil
}
