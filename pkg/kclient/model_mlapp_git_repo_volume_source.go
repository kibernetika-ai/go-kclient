/*
 * Kibernetika project, backend component
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kclient

// Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership management. Git repo volumes support SELinux relabeling.
type MlappGitRepoVolumeSource struct {
	AccessToken string `json:"access_token,omitempty"`
	AccountId   string `json:"accountId,omitempty"`
	// Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
	Directory  string `json:"directory,omitempty"`
	PrivateKey string `json:"private_key,omitempty"`
	// Repository URL
	Repository string `json:"repository"`
	// Commit hash for the specified revision.
	Revision string `json:"revision,omitempty"`
	UserName string `json:"user_name,omitempty"`
}
