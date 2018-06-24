package main

import (
	"fmt"
	"go-advanced/tree"
	"math/rand"
	"time"
)

func main() {
	BTree()
}

func BTree() {
	root := tree.New(3)
	root.Add(46)
	root.Add(11)
	root.Add(14)
	root.Add(9)
}

func BiTree() {
	var root tree.BiTree
	rand.Seed(time.Now().Unix())

	for i := 0; i < 20; i++ {
		v := rand.Int() % 100
		fmt.Printf("%d,", v)
		root.Add(v)
		//fmt.Println(root.Value())
	}
	println()
	root.PrintPreOrder()
}

func TrieNode() {
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
	fmt.Println("get age:", root.GetValue("age").Value())
}
