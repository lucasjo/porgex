package models

import (
	"time"
	//3rd party package
	"gopkg.in/mgo.v2/bson"
)

type (
	UserCapabilities struct {
		Ha                              bool     `json:"ha"`
		SubAccounts                     bool     `json:"subaccount"`
		GearSizes                       []string `json:"gear_sizes"`
		MaxDomins                       int      `json:"max_domains"`
		MaxGears                        int      `json:"max_gears"`
		MaxTeams                        int      `json:"max_teams"`
		ViewGlobalTeam                  bool     `json:"view_global_team"`
		MaxUntrackedAddtlStoragePerGear int      `json:"max_untracked_addtl_storage_per_gear"`
		MaxTrackedAddtlStoragePerGear   int      `json:"max_tracked_addtl_storage_per_gear"`
	}

	UserSshkey struct {
		ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Type    string        `json:"type"`
		_Type   string        `json:"_type"`
		Name    string        `json:"name"`
		Content string        `json:"content"`
	}

	UserPlanHistory    struct{}
	UserPendingOpGroup struct{}

	CloudUsers struct {
		ID              bson.ObjectId        `json:"id" bson:"_id,omitempty"`
		Capabilities    UserCapabilities     `json:"capabilities"`
		ConsumedGears   int                  `json:"consumed_gears"`
		CreatedAt       time.Time            `json:"create_at"`
		Login           string               `json:"login" bson:"login"`
		PendingOpGroups []UserPendingOpGroup `json:"pending_op_groups"`
		PlanHistory     []UserPlanHistory    `json:"plan_history"`
		SshKeys         []UserSshkey         `json:"ssh_keys"`
		UpdateAt        time.Time            `json:"update_at"`
	}
)
