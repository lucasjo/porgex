package models

import

//3rd party package
(
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	AnalyticsProperties struct{}
	ConfigHash          struct{}
	GroupOverride       struct{}
	MetaProperties      struct{}

	Application struct {
		ID                 bson.ObjectId       `json:"id" bson:"_id,omitempty"`
		Analytics          AnalyticsProperties `json:"analytics"`
		BuilderId          string              `json:"builder_id"`
		CanonicalName      string              `json:"canonica_name"`
		ComponentInstances []ComponentInstance `json:"component_instances"`
		Config             ConfigHash          `json:"config"`
		CreateAt           time.Time           `json:"create_at"`
		DefaultGearSize    string              `json:"default_gear_size"`
		Deployments        Deployment          `json:"deployments"`
		DomainId           bson.ObjectId       `json:"domain_id" bson:"domain_id,omitempty"`
		DomainNamespace    string              `json:"domain_namespace"`
		Gears              []Gear              `json:"gears"`
		GroupInstances     []GroupInstance     `json:"group_instances"`
		GroupOverrides     []GroupOverride     `json:"group_override"`
		Ha                 bool                `json:"ha"`
		InitGitUrl         string              `json:"init_git_url"`
		Members            []Member            `json:members"`
		Name               string              `json:"name"`
		OwnerId            bson.ObjectId       `json:"owner_id" bson:"owner_id,omitempty"`
		PendingAppOpGroups []PendingAppOpGroup `json:"pending_op_groups"`
		Scalable           bool                `json:"scalable"`
		SecretToken        string              `json:"secret_token"`
		UpdateAt           time.Time           `json:"update_at"`
		Meta               MetaProperties      `json:"meta"`
	}
)
