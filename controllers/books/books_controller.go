package books

import (
	"github.com/Narachii/crudApp/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	books, err := services.BookService.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, books)
}

func Show(c *gin.Context)  {
	id := c.Params.ByName("id")
	book, err := services.BookService.GetBookByID(id)

	if err != nil {
		c.JSON(http.StatusNoContent, err)
		return
	}
	c.JSON(http.StatusOK, book)
}

func Create(c *gin.Context)  {
	book, err := services.BookService.CreateBook(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, book)
}

func Update(c *gin.Context)  {
	id := c.Params.ByName("id")
	book, err := services.BookService.UpdateById(id, c)
	if err != nil {
		c.JSON(http.StatusNoContent, err)
		return
	}
	c.JSON(http.StatusOK, book)
}

func Delete(c *gin.Context)  {
	id := c.Params.ByName("id")
	if err := services.BookService.DeleteById(id, c); err != nil {
		c.JSON(http.StatusNoContent, err)
	}
	c.JSON(http.StatusOK, gin.H{"id #" +id: "Successfully deleted"})
}
