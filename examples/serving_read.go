package main

import (
	"github.com/kibernetika-ai/go-kclient/examples/config"
	"github.com/kibernetika-ai/go-kclient/pkg/util"
	"log"
)

func main() {
	servingWorkspaceName, servingName := "expo-recall", "demo-0-0-1"
	if serving, r, err := config.Client().ServingApi.ServingInfo(
		config.Auth(), servingWorkspaceName, servingName,
	); err != nil {
		util.ErrorFatal(err)
	} else {
		log.Printf("Status code %d, serving %s loaded", r.StatusCode, serving.Name)
	}
}
