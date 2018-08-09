package main

import (
	"fmt"
	"go-advanced/docker/compose/yamldoc"
)

func main() {
	//testYaml()
	yc := yamldoc.New()
	yc.Load("test.yaml")
	fmt.Println("version:", yc.Version)
	for k, v := range yc.Services {
		fmt.Println(k, ":")
		fmt.Println("\t", v.Image)
		fmt.Println("\t", v.ContainerName)
		fmt.Println("\t", v.Environment)
		fmt.Println("\t", v.Ports)
		fmt.Println("\t", v.Volumes)
	}
}
