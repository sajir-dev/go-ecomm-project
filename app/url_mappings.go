package app

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func mapUrls() {
	router.GET("/", ping)
	router.GET("/products", controllers.GetAllProducts)
	router.GET("/products/:id", controllers.GetProduct)
	router.POST("/products", controllers.CreateProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
}

func ping(c *gin.Context) {
	c.String(200, "Hey, I'm working...")
}
