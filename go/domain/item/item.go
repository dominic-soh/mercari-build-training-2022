package item

import (
	"mercari-build-training-2022/domain/category"
)

type Item struct {
	ID             uint              `json:"id"`
	Name           string            `json:"name"`
	Category       category.Category `json:"category"`
	Image_filename string            `json:"image_name"`
}
