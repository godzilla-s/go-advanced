package main

import (
	"fmt"
	"go-advanced/build"
	"go-advanced/build/config"
)

func main() {
	build.Add(3, 5)
	testConfig()
}

func testConfig() {
	fmt.Println(config.Host_Addr)
	fmt.Println(config.Host_Port)
	fmt.Println(config.Machine_Name)
}
