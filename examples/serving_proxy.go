package main

import (
	"github.com/kibernetika-ai/go-kclient/examples/config"
	"github.com/kibernetika-ai/go-kclient/pkg/util"
	"log"
	"time"
)

func main() {
	servingWorkspaceName, servingName, port, model := "expo-recall", "demo-0-0-1", "9000", "any"
	data := util.TFRequest(
		map[string]interface{}{
			"test_text":  "current datetime: " + time.Now().String(),
			"test_int":   42,
			"test_float": 3.1415926,
		},
		map[string]string{
			"test_image": "clover.png",
		},
	)
	if response, r, err := config.Client().ServingApi.ServingTFProxyModel(
		config.Auth(), model, data, servingWorkspaceName, servingName, port,
	); err != nil {
		util.ErrorFatal(err)
	} else {
		log.Printf("Status code %d, response: %v", r.StatusCode, response)
	}
}
