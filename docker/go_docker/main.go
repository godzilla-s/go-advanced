package main

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

func buildImages(cli *docker.Client) {
	var opt docker.BuildImageOptions
	opt.Dockerfile = "Dockerfile"

	//opt.InputStream =
	err := cli.BuildImage(opt)
	if err != nil {
		fmt.Println("build image err:", err)
		return
	}
}

func dockerInfo(cli *docker.Client) {
	info, err := cli.Info()
	if err != nil {
		fmt.Println("info error:", err)
		return
	}

	fmt.Println("ID: ", info.ID)
	fmt.Println("Images: ", info.Images)
	fmt.Println("Kernel Version: ", info.KernelVersion)
	fmt.Println("Server Version: ", info.ServerVersion)
	fmt.Println("Container: ", info.Containers)
	fmt.Println("docker network driver", info.Driver)
}

func containerList(cli *docker.Client) {
	containers, err := cli.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		fmt.Println("list container err:", err)
		return
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

}

func imageList(cli *docker.Client) {
	images, err := cli.ListImages(docker.ListImagesOptions{})
	if err != nil {
		fmt.Println("list images err:", err)
		return
	}

	for _, img := range images {
		fmt.Printf("%s | %s\n", img.ID, img.Labels)
	}
}

func runContainer(cli *docker.Client) {
	opt := docker.CreateContainerOptions{
		Name: "myhttp_server",
		Config: &docker.Config{
			Image: "myapp/httpserver",
			Cmd:   []string{"./httpserver"},
		},
		HostConfig: &docker.HostConfig{},
	}

	// 端口映射
	opt.Config.ExposedPorts = make(map[docker.Port]struct{})
	opt.Config.ExposedPorts["8010"] = struct{}{}
	opt.HostConfig.PortBindings = make(map[docker.Port][]docker.PortBinding)
	opt.HostConfig.PortBindings["8010"] = []docker.PortBinding{{HostIP: "0.0.0.0", HostPort: "8010"}}

	// volume 映射
	opt.HostConfig.Binds = []string{"tmp:/mycode"}

	ctn, err := cli.CreateContainer(opt)
	if err != nil {
		fmt.Println("create container err:", err)
		return
	}

	err = cli.StartContainer(ctn.ID, nil)
	if err != nil {
		fmt.Println("start container err:", err)
		return
	}
	fmt.Println("container start ok")
}

func stopContainer(cli *docker.Client, id string) {
	err := cli.StopContainer(id, 30)
	if err != nil {
		fmt.Println("stop container: ", id, " err:", err)
		return
	}

	opt := docker.RemoveContainerOptions{ID: id, RemoveVolumes: true}
	err = cli.RemoveContainer(opt)
	if err != nil {
		fmt.Println("remove ", id, " err:", err)
		return
	}
	fmt.Println("remove ", id, " ok")
}

func stopAllContainer(cli *docker.Client) {
	containers, err := cli.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		fmt.Println("list container:", err)
		return
	}

	for _, container := range containers {
		stopContainer(cli, container.ID)
	}
}

func main() {
	//cli, err := docker.NewClient("tcp://192.168.66.240:2375")
	cli, err := docker.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		panic(err)
	}

	//buildImages(cli)
	//dockerInfo(cli)
	//containerList(cli)
	runContainer(cli)
	//stopContainer(cli, "305b271012d6")
	//stopAllContainer(cli)
}
