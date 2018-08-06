package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func RemoteClient() {
	header := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("tcp://192.168.66.240:2375", "1.38", nil, header)
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	//containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	//fmt.Println(len(images))
	for _, image := range images {
		fmt.Printf("%s %s\n", image.ID, image.RepoTags[0])
	}
}

func LocalClient() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	//containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(len(images))
	for _, image := range images {
		fmt.Printf("%s %s\n", image.ID, image.RepoTags[0])
	}
}

func main() {
	RemoteClient()
}
