package models

import "time"

type Deployment struct {
	DeploymentId    string    `json:"deployment_id"`
	CreateAt        time.Time `json:"create_at"`
	HotDeploy       bool      `json:"hot_deploy"`
	ForceCleanBuild bool      `json:"force_clean_build"`
	Ref             string    `json:"ref"`
	Sha1            string    `json:"sha1"`
	ArtifactUrl     string    `json:"artifact_url"`
	Activations     []int     `json:"activations"`
}
