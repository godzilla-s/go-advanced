package yamldoc

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type ComposeYaml struct {
	Version  string                `yaml:"version"`
	Services map[string]composeTag `yaml:"services"`
}

type composeTag struct {
	Image         string   `yaml:"image"`
	ContainerName string   `yaml:"container_name"`
	Environment   []string `yaml:"environment"`
	Ports         []string `yaml:"ports"`
	Volumes       []string `yaml:"volumes"`
	Command       string   `yaml:"command"`
}

func New() *ComposeYaml {
	return &ComposeYaml{}
}

func (yc *ComposeYaml) Load(file string) {
	file, err := filepath.Abs("test.yaml")
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

func (cp *composeTag) Fillup() {

}
