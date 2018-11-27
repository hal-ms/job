package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var qr_host = "http://172.20.10.3:8000/qr"

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cros)
	r.GET("/alive", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/qr", func(c *gin.Context) {
		c.Redirect(http.StatusFound, qr_host)
	})
	r.Static("/_nuxt", "./spa/dist/_nuxt")
	r.LoadHTMLGlob("spa/dist/200.html")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "200.html", nil)
	})
	apiRouter(r.Group("/api"))
	return r
}

func cros(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", headers)
	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	}
	c.Set("start_time", time.Now())
	c.Next()

}
