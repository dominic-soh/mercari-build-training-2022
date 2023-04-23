package repository

import (
	"context"
	"mercari-build-training-2022/domain/item"
)

type ItemsRepository interface {
	GetAllItems(ctx context.Context) ([]item.Item, error)
	GetItemById(ctx context.Context, id string) (*item.Item, error)
	GetItemsByName(ctx context.Context, keyword string) ([]item.Item, error)
	CreateItem(ctx context.Context, name string, categoryName string, imageHash string) (*item.Item, error)
}
