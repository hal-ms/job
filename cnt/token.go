package cnt

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/cnt/util"
	"github.com/hal-ms/job/model"
	"github.com/hal-ms/job/repo"
)

func GetToken(c *gin.Context) {
	util.Json(repo.Token.Store(model.Token{}), c)
}

func AliveToken(c *gin.Context) {
	token := repo.Token.FindByID(c.Param("token"))
	if token == nil {
		util.NotFound(c)
		return
	}
	if token.IsOpen || token.Done {
		util.NotFound(c)
		return
	}
	util.NoContent(c)
}
