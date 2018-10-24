package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Job struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `json:"name"`
	DisplayName  string        `json:"display_name"`
	UserName     string        `json:"user_name"`
	Done         bool          `json:"done"`
	Notification bool          `json:"notification"`
	Created      time.Time     `json:"created"`
}
type Jobs []Job

type Token struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	IsOpen bool          `json:"is_open"`
	Done   bool          `json:"done"`
}

type JobCategory struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type JobCategorys []JobCategory
