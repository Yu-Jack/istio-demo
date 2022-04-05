package server

import (
	"example.com/product/feature"
	"github.com/gin-gonic/gin"
)

type CustomServer struct {
	Engine         *gin.Engine
	ProductService feature.ProductService
}

type CustomServerHandler interface {
	GetProductHandler(c *gin.Context)
}

func NewCutomServer(ps feature.ProductService) *CustomServer {
	engine := gin.Default()
	return &CustomServer{
		Engine:         engine,
		ProductService: ps,
	}
}
