package main

import (
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title    string
	Owner    owner
	Database database
	Server   server
	Client   client
}

type owner struct {
	Name          string
	Originization string
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type client struct {
	Data  [][]interface{}
	Hosts []string
}

func main() {
	var cfg Config
	file, err := filepath.Abs("./test.toml")
	fmt.Println(file)
	if err != nil {
		panic(err)
	}

	if _, err := toml.DecodeFile("/Users/mydev/nerthus-dev/src/go-advanced/config/toml/test.toml", &cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg.Title, cfg.Database)
}
