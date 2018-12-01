package cnt

import (
	"net/http"

	"github.com/hal-ms/job/cnt/util"

	"github.com/hal-ms/job/model"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/repo"
)

func GetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, repo.Config.Get())
}

func SetConfig(c *gin.Context) {
	var req model.Config
	err := c.BindJSON(&req)
	if err != nil {
		util.BadRequest(err.Error(), c)
		return
	}

	repo.Config.Set(req)
	util.NoContent(c)
}
