package compose

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

func testYaml() {
	file, err := filepath.Abs("test.yaml")
	if err != nil {
		panic(err)
	}
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var c ComposeYaml
	err = yaml.Unmarshal(buf, &c.data)
	if err != nil {
		panic(err)
	}

	//fmt.Println(c.data)
	for _, v := range c.data {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			vmap := v.(map[interface{}]interface{})
			for k1, v1 := range vmap {
				fmt.Println(k1)
				vtype := reflect.TypeOf(v1).Kind()
				switch vtype {
				case reflect.Map:
					v2map := v1.(map[interface{}]interface{})
					for k2, v2 := range v2map {
						vtype2 := reflect.TypeOf(v2).Kind()
						switch vtype2 {
						case reflect.Map:
							fmt.Println("\t", k2, "==> map")
						case reflect.String:
							fmt.Println("\t", k2, ":", v2)
						case reflect.Slice:
							vslice := v2.([]interface{})
							fmt.Println("\t", k2)
							for _, vs := range vslice {
								fmt.Println("\t\t-", vs)
							}

						}
					}
					//fmt.Println(v1)
				}
				//fmt.Println(k1, v1)
			}
		}

	}
}

const file = "test"

func viperTest() {
	viper.SetEnvPrefix(file)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(file)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// image := viper.GetString("service.image")
	// fmt.Println("image:", image)
	// max := viper.GetInt("peer.maxAccept")
	// fmt.Println("max:", max)

	// services := viper.GetStringSlice("service")

	services := viper.GetStringSlice("services")
	for key, srv := range services {
		//prefix := "service." + srv
		fmt.Println(key, srv)
		//fmt.Println("name: ", viper.GetString(prefix+".name"))
		//fmt.Println("image: ", viper.GetString(prefix+".image"))
	}

}
