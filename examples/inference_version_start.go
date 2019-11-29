package main

import (
	"github.com/kibernetika-ai/go-kclient/examples/config"
	"github.com/kibernetika-ai/go-kclient/pkg/kclient"
	"github.com/kibernetika-ai/go-kclient/pkg/util"
	"log"
)

func main() {
	workspaceName, inferenceName, version := "kuberlab-demo", "demo", "0.0.1"
	servingWorkspaceName, servingName := "expo-recall", "demo-0-0-1"
	req := kclient.InferenceRunServingRequest{
		WorkspaceName: servingWorkspaceName,
		Name:          servingName,
		ClusterID:     "shared/17431",
		Values: util.Pointer(map[string]interface{}{
			"params_text": "Generated in go KClient",
			"source": map[string]interface{}{
				"repository": "https://github.com/kibernetika-ai/demo-srv",
			},
		}),
	}
	if srv, r, err := config.Client().InferenceApi.InferenceInferenceVersionStart(
		config.Auth(), req, workspaceName, inferenceName, version,
	); err != nil {
		util.ErrorFatal(err)
	} else {
		log.Printf("Status code %d, serving %s started", r.StatusCode, srv.Name)
	}
}
