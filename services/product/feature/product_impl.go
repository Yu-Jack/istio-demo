package feature

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductReq struct {
	Version string `form:"version"`
}

// GetProduct only get review info of product for now.
func (rs *ProductServer) GetProduct(c *gin.Context) {
	var m ProductReq
	err := c.Bind(&m)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"messagee": "faild to bind request",
		})
		return
	}

	url := fmt.Sprintf("%s/%s", reviewEndPoint, "test")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("version", m.Version)

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{
			"messagee": "faild to get review",
		})
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	c.JSON(200, gin.H{
		"message": string(body),
	})
}
