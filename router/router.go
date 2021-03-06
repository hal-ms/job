package router

import (
	"net/http"
	"time"

	"github.com/hal-ms/job/repo"

	"github.com/gin-gonic/gin"
	iwRouter "github.com/hal-ms/job/iw/router"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cros)
	r.GET("/alive", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/qr", func(c *gin.Context) {
		c.Redirect(http.StatusFound, repo.Config.Get().Url.Main+"/qr")
	})
	r.Static("/_nuxt", "./spa/dist/_nuxt")
	r.Static("/public", "./public")
	r.LoadHTMLGlob("spa/dist/200.html")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "200.html", nil)
	})
	apiRouter(r.Group("/api"))
	iwRouter.GetRouter(r.Group("/iw"))
	return r
}

func cros(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")
	c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", headers)
	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	}
	c.Set("start_time", time.Now())
	c.Next()
}
