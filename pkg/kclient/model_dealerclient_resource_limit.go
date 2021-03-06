/*
 * Kibernetika project, backend component
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kclient

type DealerclientResourceLimit struct {
	Cpu           string `json:"cpu,omitempty"`
	CpuMi         int64  `json:"cpu_mi,omitempty"`
	ExecutionTime int64  `json:"execution_time,omitempty"`
	Gpu           int64  `json:"gpu,omitempty"`
	Memory        string `json:"memory,omitempty"`
	MemoryMb      int64  `json:"memory_mb,omitempty"`
	ParallelRuns  int64  `json:"parallel_runs,omitempty"`
	Replicas      int64  `json:"replicas,omitempty"`
}
