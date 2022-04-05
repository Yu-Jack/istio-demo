package feature

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProduct only get review info of product for now.
func (p *Product) GetProduct(c *gin.Context, payload GetProductReq) (res GetProductRes, err error) {
	url := fmt.Sprintf("%s/%s", reviewEndPoint, "test")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("version", payload.Version)

	resp, err := client.Do(req)
	if err != nil {
		return GetProductRes{}, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return GetProductRes{
		Message: string(body),
	}, nil
}
