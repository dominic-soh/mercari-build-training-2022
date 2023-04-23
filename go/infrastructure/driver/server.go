package driver

import (
	"mercari-build-training-2022/db"
	"mercari-build-training-2022/db/gateway"
	"mercari-build-training-2022/service"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Server() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.INFO)

	front_url := os.Getenv("FRONT_URL")
	if front_url == "" {
		front_url = "http://localhost:3000"
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{front_url},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	db := db.InitialiseDB()
	service := service.NewService(db, gateway.NewItemsRepository(db), gateway.NewCategoriesRepository(db))

	// Routes
	e.GET("/", service.Root)
	e.POST("/items", service.CreateItemDB)
	e.GET("/items", service.GetAllItems)
	e.GET("search", service.SearchItemByKeyword)
	e.GET("items/:itemId", service.GetItem)
	e.GET("/category", service.GetCategories)
	e.GET("/image/:imageFilename", service.GetImage)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
