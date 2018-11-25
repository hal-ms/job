package service

import (
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
