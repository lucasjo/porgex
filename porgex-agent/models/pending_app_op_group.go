package models

import "gopkg.in/mgo.v2/bson"

type PendingAppOpGroup struct {
	PendingAppOps      []PendingAppOp `json:"pending_ops"`
	ParentOpId         bson.ObjectId  `json:"parent_op_id" bson:"parent_op_id, omitempty"`
	NumGearsAdded      int            `json:"num_gears_added"`
	NumGearsRemoved    int            `json:"num_gears_removed"`
	NumGearsCreated    int            `json:"num_gears_created"`
	NumGearsRolledBack int            `json:"num_gears_roll_back"`
	UserAgent          string         `json:"user_agent"`
	RollbackBlocked    bool           `json:"rollback_blocked"`
}
