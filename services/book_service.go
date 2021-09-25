package services

import (
	"github.com/Narachii/crudApp/db"
	"github.com/Narachii/crudApp/entity"
	"github.com/gin-gonic/gin"
)

var BookService ibooksService = &booksService{}

type Book entity.Book

type booksService struct {}

type ibooksService interface {
	GetAll() ([]Book, error)
	GetBookByID(bid string) (Book, error)
	CreateBook(c *gin.Context) (Book, error)
	DeleteById(bid string, c *gin.Context) error
	UpdateById(bid string, c *gin.Context) (Book, error)
}

func (s *booksService) GetAll() ([]Book, error){
	d := db.GetDB()
	var books []Book
	if err := d.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (s *booksService) GetBookByID(bid string) (Book, error) {
	d := db.GetDB()
	var b Book

	if err := d.Where("id = ?", bid).First(&b).Error; err != nil {
		return b, err
	}
	return b, nil
}

func (s *booksService) CreateBook(c *gin.Context) (Book, error){
	d := db.GetDB()
	var b Book

	if err := c.BindJSON(&b); err != nil {
		return b, err
	}

	if err := d.Create(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

func (s *booksService) UpdateById(bid string, c *gin.Context) (Book, error) {
	d := db.GetDB()
	var b Book

	if err := d.Where("id = ?", bid).First(&b).Error; err != nil {
		return b, err
	}

	if err := c.BindJSON(&b); err != nil {
		return b, err
	}

	if err := d.Save(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

func (s *booksService) DeleteById(bid string, c *gin.Context) error {
	d := db.GetDB()
	var b Book

	if err := d.Where("id = ?", bid).Delete(&b).Error; err != nil {
		return err
	}

	return nil
}
