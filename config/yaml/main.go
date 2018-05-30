package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Product struct {
	ProductID string
	Desc      string
	Quantity  int
	Price     int
}

type Out struct {
	Id      string
	Name    string
	Product []Product
	Address string
}

func main() {
	file, err := filepath.Abs("test.yaml")
	if err != nil {
		panic(err)
	}
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var out Out
	err = yaml.Unmarshal(buf, &out)
	if err != nil {
		panic(err)
	}
	fmt.Println(out.Address, out.Name, out.Id, out.Product)
}
