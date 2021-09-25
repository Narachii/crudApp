package app

import (
	"github.com/Narachii/crudApp/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication()  {
	db.Init()
	defer db.Close()

	mapUrls()
	router.Run(":8081")
}
