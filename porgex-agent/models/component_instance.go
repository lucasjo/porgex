package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	ComponentPropertie struct{}

	ComponentInstance struct {
		ID                  bson.ObjectId      `json:"id" bson:"_id,omitempty"`
		ComponentProperties ComponentPropertie `json:"component_properties"`
		CartridgeName       string             `json:"cartridge_name"`

		ComponentName   string        `json:"component_name"`
		CartridgeVender string        `json:"cartridge_vender"`
		CartridgeId     bson.ObjectId `json:"cartridge_id" bson:"cartridge_id,omitempty"`
		GroupInstanceId bson.ObjectId `json:"group_instance_id" bson:"group_instance_id,omitempty"`
		CreateAt        time.Time     `json:"create_at"`
	}
)
