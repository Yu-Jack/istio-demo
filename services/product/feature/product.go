package feature

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
}

type ProductService interface {
	GetProduct(c *gin.Context, payload GetProductReq) (res GetProductRes, err error)
}

var reviewEndPoint string

func init() {
	reviewEndPoint = os.Getenv("REVIEW_END_POINT")
}

func NewProduct() *Product {
	return &Product{}
}
