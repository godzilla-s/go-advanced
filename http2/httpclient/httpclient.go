package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var url = ""

func main() {
	client := &http.Client{}
	req, err := client.Post(url, "application/x-www-form-urlencoded", strings.NewReader("name=zhangshan"))
	if err != nil {
		fmt.Println("err")
		return
	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
