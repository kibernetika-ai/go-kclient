package util

import (
	"github.com/kibernetika-ai/go-kclient/pkg/kclient"
	"log"
)

func ErrorFatal(err error) {
	var errMsg string
	switch errT := err.(type) {
	case kclient.GenericSwaggerError:
		errMsg = string(errT.Body())
	default:
		errMsg = errT.Error()
	}
	log.Fatal(errMsg)
}
