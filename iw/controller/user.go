package controller

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/iw/model"
	"github.com/hal-ms/job/iw/service"
)

var UserController userController

type userController struct{}

func (u *userController) Get(c *gin.Context) {
	id := c.Param("id")
	user := service.User.FindByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *userController) Create(c *gin.Context) {
	var req struct {
		IsNellow bool `json:"is_nellow"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	fmt.Println(req.IsNellow)
	var user model.User
	user.ID = bson.NewObjectId()
	user.Name = "bcp-guest-00"
	user.Icon = "hogehoge.iconUrl.com"

	err = service.User.Create(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "err user create",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *userController) Update(c *gin.Context) {
	id := c.Param("id")
	user := service.User.FindByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	var req struct {
		Name string `json:"name"`
		PId  int    `json:"p_id"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	fmt.Println(req.PId)
	if req.PId != 0 {
		user.PID = req.PId
	}

	err = service.User.Update(*user)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, nil)
}
