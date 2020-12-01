package controllers

import (
	"net/http"

	"../services"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ProductID   string  `json:"id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Discount    float32 `json:"discout"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
}

func CreateProduct(c *gin.Context) {
	var p Product
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Bad request",
			})
		return
	}

	res, err := services.ProductServices.CreateProduct(
		p.ProductName,
		p.Category,
		p.Price,
		p.Discount,
		p.Description,
		p.Rating,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		return
	}

	c.JSON(
		http.StatusOK, res,
	)

	return
}

func GetAllProducts(c *gin.Context) {
	res, err := services.ProductServices.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	res, err := services.ProductServices.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": "Internal Server Error",
			})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func UpdateProduct(c *gin.Context) {
	var p Product
	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{"error": "Bad Request"})
		return
	}

	res, err := services.ProductServices.UpdateProduct(
		p.ProductID,
		p.ProductName,
		p.Category,
		p.Price,
		p.Discount,
		p.Description,
		p.Rating,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Internal Server Error",
			})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
