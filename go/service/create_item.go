package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"mercari-build-training-2022/domain/category"
	"mercari-build-training-2022/domain/item"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type AddItemResponse struct {
	Message string    `json:"message"`
	Item    item.Item `json:"item"`
}

type AddItemRequest struct {
	Name     string                `json:"name" form:"name"`
	Category string                `json:"category" form:"category"`
	Image    *multipart.FileHeader `json:"-" form:"image"`
}

func hashImage(file []byte) string {
	hash := sha256.New()
	hash.Write([]byte(file))
	huh := hash.Sum(nil)
	extension := hex.EncodeToString(huh[:]) + ".jpg"
	return extension
}

func (s *projectService) CreateItemDB(c echo.Context) error {
	// Parse request body into AddItemRequest struct
	req := new(AddItemRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, AddItemResponse{Message: "Invalid request"})
	}
	// Get image
	image, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, AddItemResponse{Message: "No image uploaded"})
	}

	// Hash image filename
	extension := hashImage([]byte(image.Filename))
	newfile, _ := os.Create("image/" + extension)
	imgBin, err := os.ReadFile("images/" + image.Filename)
	if err != nil {
		fmt.Println(err)
	}
	newfile.Write(imgBin)

	//need to change
	newItem, err := s.itemRepo.CreateItem(c.Request().Context(), req.Name, req.Category, extension)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, AddItemResponse{Message: err.Error()})
	}

	message := fmt.Sprintf("item received: %s", image.Filename)
	return c.JSON(http.StatusOK, AddItemResponse{
		Message: message,
		Item: item.Item{
			ID:             newItem.ID,
			Name:           newItem.Name,
			Category:       category.Category{}, //change, this needs some entity
			Image_filename: newItem.Image_filename,
		},
	})
}
