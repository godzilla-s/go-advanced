package dockerclient

import (
	"fmt"
	"go-advanced/docker-compose/compose"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

type DockerClient struct {
	cli        *docker.Client
	compose    *compose.Compose
	containers ContainerSet
}

type ContainerSet map[string]*docker.Container

func New(fd string) *DockerClient {
	cli, err := docker.NewClient(fd)
	if err != nil {
		panic(err)
	}

	//var opt docker.BuildImageOptions

	//cli.BuildImage()
	return &DockerClient{
		cli:        cli,
		containers: make(ContainerSet),
	}
}

func (dc *DockerClient) LoadCompose(c *compose.Compose) {
	dc.compose = c
}

func (dc *DockerClient) CreateContainer() error {
	for k, v := range dc.compose.Services {
		opt := newContainerOpt(v)
		container, err := dc.cli.CreateContainer(opt)
		if err != nil {
			fmt.Println("create ", k, "fail", "err:", err)
			return err
		}
		//dc.cli.CreateVolume()
		fmt.Println("create ", k, "ok")
		dc.containers[k] = container
	}
	return nil
}

func newContainerOpt(s compose.Service) docker.CreateContainerOptions {
	opt := docker.CreateContainerOptions{
		Name: s.ContainerName,
		Config: &docker.Config{
			Image: s.Image,
			Env:   s.Environment, // 环境变量
			//Cmd:   strings.Split(s.Command, " "),
		},
		HostConfig: &docker.HostConfig{},
	}

	if opt.Name == "org1_fabric_ca_server" || opt.Name == "org2_fabric_ca_server" {
		fmt.Println("start fabric-ca:1.1.0")
		opt.Config.Cmd = []string{"sh", "-c", "fabric-ca-server start -b admin:adminpw -d"}
	} else {
		fmt.Println("start ", opt.Name)
		opt.Config.Cmd = strings.Split(s.Command, " ")
	}

	// volume映射
	// fmt.Println(s.Volumes)
	opt.HostConfig.Binds = s.Volumes
	// 端口映射
	opt.HostConfig.PortBindings = make(map[docker.Port][]docker.PortBinding)
	//var bindPorts []docker.PortBinding
	for _, port := range s.Ports {
		exposePorts := strings.Split(port, ":")
		if len(exposePorts) != 2 {
			panic("invalid ports set")
		}
		//bindPorts = append(bindPorts, docker.PortBinding{{HostPort: exposePorts[1]}})
		// 7053:7054  前面是主机，后面是docker， 映射是映射到主机端口
		fmt.Println("port bind ==>", docker.Port(exposePorts[1])+"/tcp", exposePorts[0])
		opt.HostConfig.PortBindings[docker.Port(exposePorts[1])+"/tcp"] = []docker.PortBinding{{HostPort: exposePorts[0] /*, HostIP: "192.168.66.240"*/}}
	}

	return opt
}

func (dc *DockerClient) Up() {
	for k, c := range dc.containers {
		fmt.Println("start container ", k, " ID:", c.ID)
		err := dc.cli.StartContainer(c.ID, nil)
		if err != nil {
			fmt.Println("start container fail:", err)
		}
	}
}

func (dc *DockerClient) Down() {
	containers, err := dc.Containers()
	if err != nil {
		fmt.Println("get container error:", err)
		return
	}

	for _, v := range containers {
		fmt.Println("status:", v.State, v.Status)

		if v.State == "running" {
			err := dc.cli.StopContainer(v.ID, 10)
			if err != nil {
				fmt.Println("stop container error:", err)
				continue
			}
		}

		var opt docker.RemoveContainerOptions
		opt.Force = true // 强制退出
		opt.ID = v.ID
		opt.RemoveVolumes = true
		err = dc.cli.RemoveContainer(opt)
		if err != nil {
			fmt.Println("remove container error:", err)
			continue
		}
		fmt.Println("remove container ", v.ID[:10], "ok")
	}
}

func (dc *DockerClient) Containers() ([]docker.APIContainers, error) {
	var opt docker.ListContainersOptions
	opt.All = true
	return dc.cli.ListContainers(opt)
}

func (dc *DockerClient) BuildImage() {
	var opt docker.BuildImageOptions

	opt.Dockerfile = ""
	opt.Name = ""
	dc.cli.BuildImage(opt)
}
