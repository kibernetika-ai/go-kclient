/*
 * Kibernetika project, backend component
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kclient

type MlappEnv struct {
	Name            string `json:"name"`
	SecretKey       string `json:"secretKey,omitempty"`
	Value           string `json:"value,omitempty"`
	ValueFromSecret string `json:"valueFromSecret,omitempty"`
}
