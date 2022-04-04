package main

import (
	"example.com/product/feature"
	"example.com/product/route"
	"github.com/gin-gonic/gin"
)

type CustomServer struct {
	Engine *gin.Engine
}

func NewCutomServer() *CustomServer {
	engine := gin.Default()
	return &CustomServer{
		Engine: engine,
	}
}

func (cs *CustomServer) Run() {
	cs.Engine.Run()
}

func main() {
	cs := NewCutomServer()
	productServer := feature.NewProductServer()
	route.Register(cs.Engine, productServer)
	cs.Run()
}
