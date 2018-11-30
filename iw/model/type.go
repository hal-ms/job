package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID                   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Icon                 string        `json:"icon"`
	Name                 string        `json:"name"`
	ProvidingDestination `json:"providing_destination" bson:"providing_destination"`
	IsNellow             bool `json:"is_nellow" bson:"is_nellow"`
	Amount               int  `json:"amount"`
}

type Users []User

type ProvidingDestination struct {
	PID   int    `json:"id"`
	PName string `json:"name"`
	PRate int    `json:"rate"`
}

type Sleeper struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID    bson.ObjectId `json:"user_id" bson:"user_id"`
	SleepTime time.Time     `json:"sleep_time" bson:"sleep_time"`
}

type NellowSleepRes struct {
	Service `json:"service"`
	KazukiUser `json:"user"`
	BedTime string `json:"bed_time"`
}

type NellowWakeupRes struct {
	User `json:"user"`
	WakeupTime string `json:"wakeup_time"`
}

type KazukiUser struct {
	ID bson.ObjectId `json:"id"`
	Icon                 string        `json:"icon"`
	Name                 string        `json:"name"`
	IsNellow             bool `json:"is_nellow" bson:"is_nellow"`
	Amount               int  `json:"amount"`
}

type Service struct {
	PID int `json:"id"`
	PRate int `json:"rate"` 
}


