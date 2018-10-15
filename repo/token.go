package repo

import (
	"github.com/hal-ms/job/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Token = tokenRepo{db.C("token")}

type tokenRepo struct {
	C *mgo.Collection
}

func (t *tokenRepo) Store(token model.Token) model.Token {
	token.ID = bson.NewObjectId()
	t.C.Insert(&token)
	return token
}
func (t *tokenRepo) Update(token model.Token) {
	t.C.UpdateId(token.ID, token)
}

func (t *tokenRepo) FindByID(id string) *model.Token {
	if !bson.IsObjectIdHex(id) {
		return nil
	}
	token := model.Token{}
	err := t.C.FindId(bson.ObjectIdHex(id)).One(&token)
	if err != nil {
		return nil
	}
	return &token
}
