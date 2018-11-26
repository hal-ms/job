package controller

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/iw/httpUtil"
	"github.com/hal-ms/job/iw/model"
	"github.com/hal-ms/job/iw/service"
	"github.com/oliamb/cutter"
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

	if req.IsNellow {
		user, err := service.User.ResetNellower()
		fmt.Println(err)
		if err != nil {
			if err != mgo.ErrNotFound {
				c.JSON(http.StatusInternalServerError, "error by reset nellower...")
				return
			}
		}
		if user != nil {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	var user model.User
	user.ID = bson.NewObjectId()
	user.Name = "bcp-guest"
	user.IsNellow = req.IsNellow
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(20) + 1
	user.Icon = "https://s3-us-west-2.amazonaws.com/dinner-match/nellow/default_img/" + strconv.Itoa(num) + ".png"
	user.PDSet(1)
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
		Name string `json:"name" binding:"max=15"`
		PId  int    `json:"p_id"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.PId != 0 {
		user.PDSet(req.PId)
	}

	err = service.User.Update(*user)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, user)
}

func (u *userController) UpdateIcon(c *gin.Context) {
	id := c.Param("id")
	user := service.User.FindByID(id)
	icon, err := c.FormFile("icon")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	file, err := icon.Open()
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "fileOpenに失敗",
		})
		return
	}

	_, format, err := image.DecodeConfig(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "画像情報取得に失敗",
		})
		return
	}

	fmt.Println(format)
	switch format {
	case "png":
	case "jpeg":
	case "gif":
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "ファイルタイプが画像じゃないですよ笑",
		})
		return
	}

	file, err = icon.Open()
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "fileOpenに失敗",
		})
		return
	}
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "画像デコードに失敗",
		})
		return
	}

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  500,
		Height: 500,
		Mode:   cutter.Centered,
	})

	name := user.ID.Hex() + "." + format
	tmpf, err := os.Create("./iw/tmp/" + name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "致命的エラー",
		})
		return
	}
	defer tmpf.Close()
	switch format {
	case "png":
		err = png.Encode(tmpf, croppedImg)
	case "jpeg":
		err = jpeg.Encode(tmpf, croppedImg, nil)
	case "gif":
		err = gif.Encode(tmpf, croppedImg, nil)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "致命的エラー | 画像保存に失敗",
		})
		return
	}

	err = httpUtil.SendFile(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "致命的エラー | 画像アップロードに失敗",
		})
		return
	}

	user.Icon = "https://s3-us-west-2.amazonaws.com/dinner-match/nellow/" + name
	service.User.Update(*user)
	c.JSON(http.StatusOK, user)

}
