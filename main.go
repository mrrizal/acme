package main

import (
	"acme/internal/config"
	"acme/internal/rest"

	"context"
)

func main() {
	// bind stop channel to context
	ctx := context.Background()

	// start REST server
	server := rest.New(config.App.Address)
	server.Listen(ctx.Done())
}
