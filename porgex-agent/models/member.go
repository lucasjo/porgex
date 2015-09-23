package models

import "gopkg.in/mgo.v2/bson"

type Member struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Type         string        `json:"type"`
	Name         string        `json:"Name"`
	From         []string      `json:"from"`
	Role         string        `json:"role"`
	ExplicitRole string        `json:"explicit_role"`
}
