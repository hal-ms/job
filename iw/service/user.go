package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hal-ms/job/iw/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var User = userService{db.C("users")}

type userService struct {
	C *mgo.Collection
}

func (u *userService) FindAll() (model.Users, error) {
	var users model.Users
	err := u.C.Find(nil).All(&users)
	return users, err
}

func (u *userService) FindByID(id string) *model.User {
	var user model.User
	if !bson.IsObjectIdHex(id) {
		return nil
	}

	err := u.C.FindId(bson.ObjectIdHex(id)).One(&user)
	if err == mgo.ErrNotFound {
		return nil
	}
	return &user
}

func (u *userService) ResetNellower() (*model.User, error) {
	var user model.User
	err := u.C.Find(bson.M{"is_nellow": true}).One(&user)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	user.Name = "bcp-guest"
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(20) + 1
	user.Icon = "https://s3-us-west-2.amazonaws.com/dinner-match/nellow/default_img/" + strconv.Itoa(num) + ".png"
	user.PDSet(1)
	err = u.Update(user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) Create(user model.User) error {
	err := u.C.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) Update(user model.User) error {
	return u.C.UpdateId(user.ID, user)
}
