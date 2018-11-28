package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID                   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Icon                 string        `json:"icon"`
	Name                 string        `json:"name"`
	ProvidingDestination `json:"providing_destination" bson:"providing_destination"`
	IsNellow             bool 			`json:"is_nellow" bson:"is_nellow"`
	Amount				 int			`json:"amount"`
}

type Users []User

type ProvidingDestination struct {
	PID   int    `json:"id"`
	PName string `json:"name"`
	PRate uint   `json:"rate"`
}
