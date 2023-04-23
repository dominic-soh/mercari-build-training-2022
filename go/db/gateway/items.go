package gateway

import (
	"context"
	"errors"
	"mercari-build-training-2022/db/models"
	"mercari-build-training-2022/domain/category"
	"mercari-build-training-2022/domain/item"

	"gorm.io/gorm"
)

type ItemsRepository struct {
	client *gorm.DB
}

func NewItemsRepository(client *gorm.DB) *ItemsRepository {
	return &ItemsRepository{
		client: client,
	}
}

func (ir *ItemsRepository) GetAllItems(ctx context.Context) ([]item.Item, error) {
	var items []item.Item
	results := ir.client.Table("items").
		Select("items.*, categories.name as category_name").
		Joins("left join categories on items.category_id = categories.id").
		Scan(&items)
	if results.Error != nil {
		return nil, results.Error
	}

	return items, nil
}

func (ir *ItemsRepository) GetItemById(ctx context.Context, id string) (*item.Item, error) {
	var item *item.Item
	results := ir.client.Table("items").
		Select("items.*, categories.name as category_name").
		Joins("left join categories on items.category_id = categories.id").
		Where("items.id = ?", id).
		Find(&item)
	if results.Error != nil {
		return nil, results.Error
	}

	return item, nil
}

func (ir *ItemsRepository) GetItemsByName(ctx context.Context, keyword string) ([]item.Item, error) {
	var items []item.Item
	results := ir.client.Table("items").
		Select("items.*, categories.name as category_name").
		Joins("left join categories on items.category_id = categories.id").
		Where("items.name = ?", keyword).
		Find(&items)
	if results.Error != nil {
		return nil, results.Error
	}

	return items, nil
}

//use item.Item as the input, then return model as example
func (ir *ItemsRepository) CreateItem(ctx context.Context, name string, categoryName string, imageHash string) (*item.Item, error) {
	var newItem models.Item
	var cat models.Category

	// Start a transaction
	tx := ir.client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check if the category already exists
	result := tx.Where("name = ?", categoryName).First(&cat)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, result.Error
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Create a new category if it doesn't exist
		cat.Name = categoryName
		result = tx.Create(&cat)
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	// Create the new item
	newItem.Name = name
	newItem.CategoryID = cat.ID
	newItem.Image_filename = imageHash
	result = tx.Create(&newItem)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// Commit the transaction
	result = tx.Commit()
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	return &item.Item{
		ID:   newItem.ID,
		Name: newItem.Name,
		Category: category.Category{
			ID:   cat.ID,
			Name: cat.Name,
		},
		Image_filename: newItem.Image_filename,
	}, nil
}
