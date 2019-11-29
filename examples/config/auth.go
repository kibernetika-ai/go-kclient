package config

import (
	"context"
	"github.com/kibernetika-ai/go-kclient/pkg/kclient"
	"os"
)

var ctx context.Context

func Auth() context.Context {
	if ctx == nil {
		ctx = context.WithValue(context.Background(), kclient.ContextAPIKey, kclient.APIKey{
			Key:    os.Getenv("KCLIENT_AUTH"),
			Prefix: "Bearer", // Omit if not necessary.
		})
	}
	return ctx
}
