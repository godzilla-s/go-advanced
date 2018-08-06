package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("new client error:", err)
		return
	}

	ctx := context.Background()
	cli.ContainerCreate(ctx, &container.Config{
		Image: "",
		Cmd:   []string{""},
	}, nil, nil, "httpserver")
}
