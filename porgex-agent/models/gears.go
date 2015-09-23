package models

import "gopkg.in/mgo.v2/bson"

type Gear struct {
	ServerIdentity  string        `json:"server_identity"`
	Uuid            string        `json:"uuid"`
	Uid             int           `json:"uid"`
	Name            string        `json:"name"`
	Quarntined      bool          `json:"quarntined"`
	Removed         bool          `json:"removed"`
	HostSingletons  bool          `json:"host_singletons"`
	AppDns          bool          `json:"app_dns"`
	SparseCarts     []string      `json:"sparse_carts"`
	GroupInstanceId bson.ObjectId `json:"group_instance_id" bson:"group_instance_id,omitempty"`
	PortInterfaces  PortInterface `json:"port_interface"`
}
