package main

import (
	"github.com/kibernetika-ai/go-kclient/examples/config"
	"github.com/kibernetika-ai/go-kclient/pkg/util"
	"log"
)

func main() {
	workspaceName, inferenceName, version := "kuberlab-demo", "demo", "0.0.1"
	if ver, r, err := config.Client().InferenceApi.InferenceInferenceVersionInfo(
		config.Auth(), workspaceName, inferenceName, version,
	); err != nil {
		util.ErrorFatal(err)
	} else {
		log.Printf("Status code %d, version %s loaded", r.StatusCode, ver.Version)
	}
}
