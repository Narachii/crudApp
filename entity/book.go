package entity

type Book struct {
	Id uint `gorm:"primary_key" json:"id"`
	Author string `json:"author"`
	Name string `json:"name"`
	PageCount int `json:"page_count"`
}
