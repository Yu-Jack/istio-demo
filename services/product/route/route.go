package route

import (
	"github.com/gin-gonic/gin"
)

type Route interface {
	GetProduct(c *gin.Context)
}

func Register(c *gin.Engine, route Route) {
	c.GET("/product", route.GetProduct)
}
