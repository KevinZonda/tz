package main

import (
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"

	"github.com/KevinZonda/tz/handler"
)

func main() {
	g := gin.New()
	g.GET("/", func(c *gin.Context) {
		c.String(200, handler.Status())
	})
	g.GET("/nv", func(c *gin.Context) {
		bs, err := exec.Command("nvidia-smi").Output()
		if err != nil {
			c.String(500, "NV SMI is not available.")
		}
		c.String(200, string(bs))
	})
	g.Run(os.Args[1])
}
