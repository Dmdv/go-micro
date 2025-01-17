package web

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) error {
	files, err := WalkDirs("", false)
	if err != nil {
		return err
	}
	for _, f := range files {
		router.GET(f, func(name string) gin.HandlerFunc {
			return func(c *gin.Context) {
				data, err := ReadFile(name)
				if err != nil {
					c.AbortWithError(500, err)
					return
				}
				switch filepath.Ext(name) {
				case ".html":
					c.Header("Content-Type", "text/html; charset=utf-8")
				case ".css":
					c.Header("Content-Type", "text/css; charset=utf-8")
				case ".js":
					c.Header("Content-Type", "application/javascript")
				case ".svg":
					c.Header("Content-Type", "image/svg+xml")
				}
				if _, err := c.Writer.Write(data); err != nil {
					c.AbortWithError(500, err)
					return
				}
			}
		}(f))
	}
	router.GET("/", func(c *gin.Context) {
		data, err := ReadFile("index.html")
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.Header("Content-Type", "text/html; charset=utf-8")
		if _, err := c.Writer.Write(data); err != nil {
			c.AbortWithError(500, err)
			return
		}
	})
	return nil
}
