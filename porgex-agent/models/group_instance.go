package models

type GroupInstance struct {
	PlatForm  string `json:"platform"`
	GearSize  string `json:"gear_size"`
	AddtlFsGb int    `json:"addtl_fs_gb"`
}
