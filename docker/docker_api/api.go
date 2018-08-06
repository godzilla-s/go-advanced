package docker_api

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

type DockerClient struct {
	cli *docker.Client
}

func New(fd string) *DockerClient {
	cli, err := docker.NewClient(fd)
	if err != nil {
		fmt.Println("err:", err)
		return nil
	}

	return &DockerClient{
		cli: cli,
	}
}

func (dc *DockerClient) DockerInfo() (*docker.DockerInfo, error) {
	info, err := dc.cli.Info()
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (dc *DockerClient) ImagesList() {
	dc.cli.ListImages(docker.ListImagesOptions{})
}
