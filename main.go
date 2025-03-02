package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"

	"github.com/KevinZonda/tz/handler"
)

func wrap(f func() string) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshInterval := 5 // Default refresh interval in seconds
		html := `<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="refresh" content="` + fmt.Sprint(refreshInterval) + `">
    <title>System Monitor</title>
</head>
<body>
    <pre>` + f() + `</pre>
</body>
</html>`
		c.Header("Content-Type", "text/html")
		c.String(200, html)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.GET("/", wrap(handler.Status))
	g.GET("/cpu", wrap(handler.Cpu))
	g.GET("/mem", wrap(handler.Mem))
	g.GET("/nv", func(c *gin.Context) {
		bs, err := exec.Command("nvidia-smi").Output()
		if err != nil {
			c.String(500, "NV SMI is not available.")
		}
		c.String(200, string(bs))
	})
	g.Run(os.Args[1])
}
