/*
 * Kibernetika project, backend component
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kclient

type MlappDatasetSource struct {
	Dataset   string `json:"dataset,omitempty"`
	ServerURL string `json:"serverURL,omitempty"`
	Version   string `json:"version,omitempty"`
	Workspace string `json:"workspace"`
}