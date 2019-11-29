package main

import (
	"github.com/kibernetika-ai/go-kclient/examples/config"
	"github.com/kibernetika-ai/go-kclient/pkg/util"
	"log"
)

func main() {
	servingWorkspaceName, servingName := "expo-recall", "demo-0-0-1"
	if r, err := config.Client().ServingApi.ServingDelete(
		config.Auth(), servingWorkspaceName, servingName, nil,
	); err != nil {
		util.ErrorFatal(err)
	} else {
		log.Printf("Status %d, serving successfully deleted", r.StatusCode)
	}
}
