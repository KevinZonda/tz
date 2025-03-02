package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/KevinZonda/tz/handler"
)

func wrap(f func() string) gin.HandlerFunc {
	return func(c *gin.Context) {
		refQuery := c.Query("refresh")
		if refQuery == "" {
			refQuery = "1"
		}

		refreshInterval, err := strconv.Atoi(refQuery)
		if err != nil {
			c.String(400, "Invalid refresh interval")
			return
		}

		var refreshMeta string
		if refreshInterval > 0 {
			refreshMeta = `<meta http-equiv="refresh" content="` + fmt.Sprint(refreshInterval) + `">`
		}

		html := `<!DOCTYPE html>
<html>
<head>
    ` + refreshMeta + `
    <title>System Monitor</title>
</head>
<body>
    <pre>` + f() +
			"\n\nRefresh interval: " + fmt.Sprint(refreshInterval) + " seconds" +
			"\nRefreshed at: " + time.Now().Format(time.RFC3339) +
			`</pre>
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
