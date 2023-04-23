package service

import (
	"fmt"
	"mercari-build-training-2022/domain/item"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	ImgDir = "image"
)

type GetImageResponse struct {
	Message string      `json:"message"`
	Items   []item.Item `json:"items"`
}

func (s *projectService) GetImage(c echo.Context) error {
	// Create image path
	imgPath := path.Join(ImgDir, c.Param("imageFilename"))

	if !strings.HasSuffix(imgPath, ".jpg") {
		res := GetImageResponse{Message: "Image path does not end with .jpg"}
		return c.JSON(http.StatusBadRequest, res)
	}
	if _, err := os.Stat(imgPath); err != nil {
		fmt.Println(err)
		c.Logger().Debugf("Image not found: %s", imgPath)
		imgPath = path.Join(ImgDir, "default.jpg")
	}
	return c.File(imgPath)
}
