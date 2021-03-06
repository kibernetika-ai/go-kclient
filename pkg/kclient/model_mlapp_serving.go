/*
 * Kibernetika project, backend component
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kclient

type MlappServing struct {
	Args                 string                `json:"args,omitempty"`
	Autoscale            *MlappAutoscale       `json:"autoscale,omitempty"`
	Build                string                `json:"build,omitempty"`
	BuildInfo            *interface{}          `json:"build_info,omitempty"`
	Command              string                `json:"command,omitempty"`
	DefaultMountPath     string                `json:"default_mount_path,omitempty"`
	DefaultVolumeMapping bool                  `json:"default_volume_mapping,omitempty"`
	Disabled             bool                  `json:"disabled"`
	DisplayName          string                `json:"displayName,omitempty"`
	Env                  []MlappEnv            `json:"env,omitempty"`
	FrontApi             string                `json:"front_api,omitempty"`
	Images               *MlappImages          `json:"images,omitempty"`
	Labels               map[string]string     `json:"labels,omitempty"`
	Name                 string                `json:"name,omitempty"`
	Nodes                string                `json:"nodes,omitempty"`
	Ports                []MlappPort           `json:"ports,omitempty"`
	Replicas             int32                 `json:"replicas,omitempty"`
	Resources            *MlappResourceRequest `json:"resources,omitempty"`
	SkipPrefix           bool                  `json:"skipPrefix"`
	Spec                 *MlappServingSpec     `json:"spec,omitempty"`
	TaskName             string                `json:"taskName,omitempty"`
	Volumes              []MlappVolumeMount    `json:"volumes,omitempty"`
	WorkDir              string                `json:"workDir,omitempty"`
}
