package feature

import (
	"fmt"
	"os"
)

var reviewEndPoint string

func init() {
	reviewEndPoint = fmt.Sprintf("http://%s", os.Getenv("REVIEW_END_POINT"))
}

type ProductServer struct {
}

func NewProductServer() *ProductServer {
	return &ProductServer{}
}
