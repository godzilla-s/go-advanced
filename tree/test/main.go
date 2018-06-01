package main

import (
	"fmt"
	"go-advanced/tree"
)

func main() {
	root := new(tree.TrieNode)
	root.Add("hello")
	root.Add("helpp")
	root.Peek()
	root.Add("ding")
	root.Add("dong")
	root.Add("paris")
	root.Peek()
	fmt.Println(root.Find("helu"))

	root.AddValue("name", "zhangshan")
	root.AddValue("age", 230)

	fmt.Println("get name:", root.GetValue("name").Value())
}
