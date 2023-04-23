package repository

import (
	"context"
	"mercari-build-training-2022/domain/category"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) ([]category.Category, error)
	CreateCategory(ctx context.Context, name string) (*category.Category, error)
	GetCategory(ctx context.Context, name string) (*category.Category, error)
}
