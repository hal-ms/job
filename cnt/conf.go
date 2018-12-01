package cnt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/repo"
)

func GetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, repo.Config.Get())
}
