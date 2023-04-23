package models

type Category struct {
	ID   uint
	Name string `gorm:"unique"`
}
