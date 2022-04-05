package server

import (
	"fmt"
	"net/http"

	"example.com/product/feature"
	"github.com/gin-gonic/gin"
)

func (cs *CustomServer) Run() {
	cs.Engine.Run()
}

// RegisterRoutes register all routes.
func (cs *CustomServer) RegisterRoutes() {
	cs.Engine.GET("/product", cs.GetProductHandler)
}

// GetProductHandler is layout for handling the request.
func (cs *CustomServer) GetProductHandler(c *gin.Context) {
	var payload feature.GetProductReq
	err := c.Bind(&payload)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"messagee": "faild to bind request",
		})
		return
	}

	res, err := cs.ProductService.GetProduct(c, payload)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "faild to get product",
		})
	}

	c.JSON(http.StatusOK, res)
}
