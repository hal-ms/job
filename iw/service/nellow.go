package service

import (
	"github.com/hal-ms/job/iw/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var Nellow = nellowService{db.C("sleepers")}

type nellowService struct {
	C *mgo.Collection
}

func (n *nellowService) Create(id bson.ObjectId, now time.Time) error {
	var ns = model.Sleeper{
		UserID:    id,
		SleepTime: now,
	}
	return n.C.Insert(ns)
}

func (n *nellowService) FindByUserID(id bson.ObjectId) *model.Sleeper {
	var ns model.Sleeper
	err := n.C.Find(bson.M{"user_id": id}).One(&ns)
	if err != nil {
		return nil
	}
	return &ns
}

func (n *nellowService) Delete(ns model.Sleeper) error {
	return n.C.RemoveId(ns.ID)
}
