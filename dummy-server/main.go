package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ResponseMockDTO struct {
	Code     int    `json:"code"`
	Body     string `json:"body"`
	RespTime int64  `json:"resp_time"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/delay", func(c *gin.Context) {
		var req ResponseMockDTO
		if err := c.ShouldBind(&req); err != nil {

		}
		if req.RespTime > 0 {
			time.Sleep(time.Duration(req.RespTime) * time.Second)
		}
		var code int
		if req.Code != 0 {
			code = req.Code
		} else {
			code = http.StatusOK
		}

		c.JSON(code, req.Body)
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(); err != nil {
		panic(err)
	}

}
