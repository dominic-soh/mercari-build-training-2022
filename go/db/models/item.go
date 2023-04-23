package models

type Item struct {
	ID             uint `json:"-"`
	Name           string
	CategoryID     uint `gorm:"index;not null;ForeignKey:CategoryID"`
	Image_filename string
}
