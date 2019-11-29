package config

import (
	"github.com/kibernetika-ai/go-kclient/pkg/kclient"
)

var client *kclient.APIClient

func Client() *kclient.APIClient {
	if client == nil {
		client = kclient.NewAPIClient(kclient.NewConfiguration())
	}
	return client
}
