package util

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BadRequest(err string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"err": err,
	})
}
func Status(status int, c *gin.Context) {
	start, ok := c.Get("start_time")
	if ok {
		c.Header("X-Server-Latency", time.Since(start.(time.Time)).String())
	}
	c.Status(status)
	c.Abort()
}

func Created(c *gin.Context) {
	Status(http.StatusCreated, c)
}
func NoContent(c *gin.Context) {
	Status(http.StatusNoContent, c)
}
func NotFound(c *gin.Context) {
	Status(http.StatusNotFound, c)
}

func Json(obj interface{}, c *gin.Context) {
	start, ok := c.Get("start_time")
	if ok {
		c.Header("X-Server-Latency", time.Since(start.(time.Time)).String())
	}

	c.JSON(http.StatusOK, obj)
}
