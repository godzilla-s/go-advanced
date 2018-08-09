package compose

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Compose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

// compose service文件
type Service struct {
	Image         string   `yaml:"image"`
	ContainerName string   `yaml:"container_name"`
	Environment   []string `yaml:"environment"`
	Ports         []string `yaml:"ports"`
	Volumes       []string `yaml:"volumes"`
	Command       string   `yaml:"command"`
	DependsOn     []string `yaml:"depends_on"`
	WorkingDir    string   `yaml:"working_dir"`
}

func New() *Compose {
	return &Compose{}
}

func (yc *Compose) Load(filename string) {
	file, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	//var c ComposeYaml
	err = yaml.Unmarshal(buf, yc)
	if err != nil {
		panic(err)
	}
}

func (srv *Service) Fillup() {
}
