package main

import (
	"example.com/product/feature"
	"example.com/product/server"
)

func main() {
	product := feature.NewProduct()
	cs := server.NewCutomServer(product)
	cs.RegisterRoutes()
	cs.Run()
}
