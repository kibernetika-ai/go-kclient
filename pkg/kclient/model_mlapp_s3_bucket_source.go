/*
 * Kibernetika project, backend component
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kclient

type MlappS3BucketSource struct {
	AccountId string `json:"accountId,omitempty"`
	Bucket    string `json:"bucket"`
	Region    string `json:"region,omitempty"`
	Server    string `json:"server,omitempty"`
}
