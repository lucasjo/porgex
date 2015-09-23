package models

import "gopkg.in/mgo.v2/bson"

type PendingAppOp struct {
	State           string        `json:"state"`
	Prereq          []string      `json:"prereq"`
	RetryCount      int           `json:"retry_count"`
	RetryRollbackOp bson.ObjectId `json:"retry_rollback_op" bson:"retry_rollback_op,omitempty"`
	SkipRollback    bool          `json:"skip_rollback"`
}
