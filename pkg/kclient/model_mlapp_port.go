/*
 * Kibernetika project, backend component
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kclient

type MlappPort struct {
	Name       string `json:"name"`
	Port       int32  `json:"port,omitempty"`
	Protocol   string `json:"protocol,omitempty"`
	TargetPort int32  `json:"targetPort,omitempty"`
}
