// 字典树

package tree

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	endTag bool
	child  int
	value  interface{}
	Nodes  [26]*TrieNode
}

// Add 添加节点
func (tn *TrieNode) Add(prefix string) {
	bytes := []byte(strings.ToLower(prefix))
	root := tn
	for _, c := range bytes {
		if root.Nodes[c-97] == nil {
			root.Nodes[c-97] = new(TrieNode)
			root.child++
			root = root.Nodes[c-97]
		} else {
			root.child++
			root = root.Nodes[c-97]
		}
	}
	root.endTag = true
}

// AddValue
func (tn *TrieNode) AddValue(key string, val interface{}) {
	bytes := []byte(strings.ToLower(key))
	root := tn
	for _, c := range bytes {
		if root.Nodes[c-97] == nil {
			root.Nodes[c-97] = new(TrieNode)
			root.child++
			root = root.Nodes[c-97]
		} else {
			root.child++
			root = root.Nodes[c-97]
		}
	}
	root.endTag = true
	root.value = val
}

func (tn *TrieNode) GetValue(key string) *TrieNode {
	return tn.find(key)
}

// Peek 查看数据
func (tn *TrieNode) Peek() {
	i := 0
	root := tn
	for _, n := range root.Nodes {
		if n != nil {
			fmt.Println(i, n.child)
			root.Nodes[i].Peek()
			if n.endTag {
				fmt.Println("---- end ---")
			}
		}
		i++
	}
}

func (tn *TrieNode) find(key string) *TrieNode {
	bytes := []byte(strings.ToLower(key))
	root := tn
	for _, c := range bytes {
		root = root.Nodes[c-97]
		if root == nil {
			return nil
		}
	}
	if root.endTag {
		return root
	}
	return nil
}

func (tn *TrieNode) Value() interface{} {
	return tn.value
}

// Find 查找前缀
func (tn *TrieNode) Find(prefix string) int {
	bytes := []byte(strings.ToLower(prefix))
	root := tn
	i := 0
	for _, c := range bytes {
		root = root.Nodes[c-97]
		if root == nil {
			return i
		}
		i++
	}
	return i
}

// Delete 删除
func (tn *TrieNode) Delete(prefix string) {

}
