package main

import (
	"flag"
	"go-advanced/docker-compose/compose"
	"go-advanced/docker-compose/dockerclient"
)

var (
	up   bool
	down bool
)

func init() {
	flag.BoolVar(&up, "up", false, "start containers")
	flag.BoolVar(&down, "down", false, "stop containers")
	flag.Parse()
}

func main() {
	docker := dockerclient.New("unix:///var/run/docker.sock")
	//fmt.Println(up, down)
	if up {
		c := compose.New()
		c.Load("docker-cli.yaml")
		//fmt.Println("load file ok")
		docker.LoadCompose(c)
		err := docker.CreateContainer()
		if err != nil {
			return
		}
		docker.Up()
	}

	if down {
		docker.Down()
	}
}
