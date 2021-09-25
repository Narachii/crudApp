package app

import "github.com/Narachii/crudApp/controllers/books"

func mapUrls()  {
	b := router.Group("/books")
	{
		b.GET("", books.Index)
		b.GET("/:id", books.Show)
		b.POST("", books.Create)
		b.PUT("/:id", books.Update)
		b.DELETE("/:id", books.Delete)
	}
}
