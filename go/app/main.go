package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	ImgDir = "image"
)

type Response struct {
	Message string `json:"message"`
}

type ItemResponse struct {
	ID       uint `json:"-"`
	Name     string
	Category string
	Image    string
}

type Item struct {
	ID         uint `json:"-"`
	Name       string
	CategoryID uint
	Image      string
}

type Category struct {
	ID   uint
	Name string `gorm:"unique"`
}

type ItemsResponseArray struct {
	Items []ItemResponse `json:"items"`
}

func root(c echo.Context) error {
	res := Response{Message: "Hello, world!"}
	return c.JSON(http.StatusOK, res)
}

// func addItem(c echo.Context) error {
// 	// Get form data
// 	name := c.FormValue("name")
// 	category := c.FormValue("category")
// 	c.Logger().Infof("Receive item: %s", name)
// 	c.Logger().Infof("Receive item: %s", category)
// 	itemised := itemise(name, category)
// 	// Create File
// 	file, err := os.OpenFile("items.json", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
// 	if err != nil {
// 		fmt.Println(err)
// 		filebytes, e := os.ReadFile("items.json")
// 		if e != nil {
// 			fmt.Println(e)
// 		}
// 		// Convert to str
// 		existingItems := string(filebytes)
// 		fmt.Println(existingItems)
// 		jsonExistingItems := decode(existingItems)
// 		fmt.Println(jsonExistingItems)
// 		jsonExistingItems.addToItemArray(itemised)
// 		fmt.Println(jsonExistingItems)
// 		jsonData, error := json.Marshal(jsonExistingItems)
// 		if error != nil {
// 			fmt.Println(error)
// 		}
// 		newfile, _ := os.Create("items.json")
// 		newfile.Write(jsonData)
// 	} else {
// 		jsonData := appendItem(itemised)
// 		file.Write(jsonData)
// 	}

// 	message := fmt.Sprintf("item received: %s", name)
// 	res := Response{Message: message}

// 	return c.JSON(http.StatusOK, res)
// }

// func appendItem(itemised Item) []byte {
// 	items := []Item{}
// 	itemsArray := ItemsArray{items}
// 	itemsArray.addToItemArray(itemised)
// 	jsonData, err := json.Marshal(itemsArray)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return jsonData
// }

// func decode(jsonString string) ItemsArray {
// 	var stcData ItemsArray
// 	if err := json.Unmarshal([]byte(jsonString), &stcData); err != nil {
// 		fmt.Println(err)
// 	}
// 	return stcData
// }

// func itemise(name string, category string) Item {
// 	item := Item{Name: name, Category: category}
// 	return item
// }

// func (itemsArray *ItemsArray) addToItemArray(item Item) []Item {
// 	itemsArray.Items = append(itemsArray.Items, item)
// 	return itemsArray.Items
// }

// func getItems(c echo.Context) error {
// 	filebytes, e := os.ReadFile("items.json")
// 	var msg ItemsArray
// 	if e != nil {
// 		fmt.Println(e)

// 	}
// 	err := json.Unmarshal(filebytes, &msg)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return c.JSON(http.StatusOK, msg)
// }

func getImg(c echo.Context) error {
	// Create image path
	imgPath := path.Join(ImgDir, c.Param("itemImg"))

	if !strings.HasSuffix(imgPath, ".jpg") {
		res := Response{Message: "Image path does not end with .jpg"}
		return c.JSON(http.StatusBadRequest, res)
	}
	if _, err := os.Stat(imgPath); err != nil {
		c.Logger().Debugf("Image not found: %s", imgPath)
		imgPath = path.Join(ImgDir, "default.jpg")
	}
	return c.File(imgPath)
}

func addItemDB(c echo.Context) error {
	// Get form data
	name := c.FormValue("name")
	category := c.FormValue("category")
	image := c.FormValue("image")
	// Initialise DB
	db := initialiseDB()
	db.Create(&Category{Name: category})
	// Find CategoryID
	var categoryDBObj Category
	db.Where("name = ?", category).First(&categoryDBObj)

	// Hash image
	file, err := os.ReadFile("./images/" + image)
	if err != nil {
		fmt.Println(err)
		c.Logger().Debugf("Image not found: %s", image)
		imgPath := path.Join(ImgDir, "default.jpg")
		defaultFile, _ := os.ReadFile(imgPath)
		extension := hashImage(defaultFile)
		// Create
		db.Create(&Item{Name: name, CategoryID: categoryDBObj.ID, Image: extension})
		message := fmt.Sprintf("item received: %s", name)
		res := Response{Message: message}

		return c.JSON(http.StatusOK, res)
	} else {
		extension := hashImage(file)

		// Create
		db.Create(&Item{Name: name, CategoryID: categoryDBObj.ID, Image: extension})

		message := fmt.Sprintf("item received: %s", name)
		res := Response{Message: message}

		return c.JSON(http.StatusOK, res)
	}
}

func initialiseDB() *gorm.DB {
	// Initialise DB
	db, err := gorm.Open(sqlite.Open("../db/items.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Category{})

	return db
}

func hashImage(file []byte) string {
	hash := sha256.New()
	hash.Write([]byte(file))
	huh := hash.Sum(nil)
	extension := hex.EncodeToString(huh[:]) + ".jpg"
	return extension
}

func getItemsDB(c echo.Context) error {
	// Initialise DB
	db := initialiseDB()

	// Read
	var itemsResponse []ItemResponse
	db.Table("items").Select("items.id", "items.name as name", "categories.name as category", "items.image").Joins("left join categories on categories.id = items.category_id").Find(&itemsResponse)

	itemsResponseCopy := make([]ItemResponse, len(itemsResponse))
	copy(itemsResponseCopy, itemsResponse)
	var itemsResponseArray = ItemsResponseArray{itemsResponseCopy}

	return c.JSON(http.StatusOK, itemsResponseArray)
}

func getCategoryDB(c echo.Context) error {
	// Initialise DB
	db := initialiseDB()

	// Read
	var category []Category
	db.Find(&category)

	return c.JSON(http.StatusOK, category)
}

func getItemDetailDB(c echo.Context) error {
	// Get ID
	id := c.Param("itemId")

	// Initialise DB
	db := initialiseDB()

	// Read
	var itemResponse ItemResponse
	db.Table("items").Select("items.id", "items.name as name", "categories.name as category", "items.image").Joins("left join categories on categories.id = items.category_id").Where("items.id = ?", id).Find(&itemResponse)
	return c.JSON(http.StatusOK, itemResponse)
}

// func getItems(c echo.Context) error {
// 	filebytes, e := os.ReadFile("items.json")
// 	var msg ItemsArray
// 	if e != nil {
// 		fmt.Println(e)

// 	}
// 	err := json.Unmarshal(filebytes, &msg)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return c.JSON(http.StatusOK, msg)
// }

func searchItems(c echo.Context) error {
	// Get form values
	keyword := c.QueryParam("keyword")

	// Initialise DB
	db := initialiseDB()

	// Search
	var items []Item
	db.Where("name = ?", keyword).Find(&items)
	var itemsResponse []ItemResponse
	db.Table("items").Select("items.id", "items.name as name", "categories.name as category", "items.image").Joins("left join categories on categories.id = items.category_id").Where("items.name = ?", keyword).Find(&itemsResponse)

	itemsResponseCopy := make([]ItemResponse, len(itemsResponse))
	copy(itemsResponseCopy, itemsResponse)
	var itemsResponseArray = ItemsResponseArray{itemsResponseCopy}

	return c.JSON(http.StatusOK, itemsResponseArray)
}

func main() {
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

	// Routes
	e.GET("/", root)
	e.POST("/items", addItemDB)
	e.GET("/items", getItemsDB)
	e.GET("search", searchItems)
	e.GET("/image/:itemImg", getImg)
	e.GET("items/:itemId", getItemDetailDB)
	e.GET("/category", getCategoryDB)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
