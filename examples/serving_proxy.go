package main

import (
	"github.com/kibernetika-ai/go-kclient/examples/config"
	"github.com/kibernetika-ai/go-kclient/pkg/util"
	"log"
	"time"
)

func main() {
	servingWorkspaceName, servingName := "expo-recall", "demo-0-0-1"
	data := util.Request(
		map[string]interface{}{
			"test_text":  "current datetime: " + time.Now().String(),
			"test_int":   42,
			"test_float": 3.1415926,
		},
		map[string]string{
			"test_image": "clover.png",
		},
	)
	if response, r, err := config.Client().ServingApi.ServingProxy(
		config.Auth(), data, servingWorkspaceName, servingName,
	); err != nil {
		util.ErrorFatal(err)
	} else {
		log.Printf("Status code %d, response: %v", r.StatusCode, response)
	}
}
