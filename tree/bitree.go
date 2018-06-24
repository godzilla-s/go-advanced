package tree

import "fmt"

// 二叉树
type node struct {
	parent *node
	left   *node
	right  *node
	val    int
}

func newNode(val int) *node {
	return &node{
		val: val,
	}
}

type BiTree struct {
	root *node
}

func (b *BiTree) Value() int {
	return b.root.val
}

func (b *BiTree) Add(val int) {
	if b.root == nil {
		b.root = newNode(val)
		return
	}

	node := b.root
	//fmt.Println(node.val, val)
	for node != nil {
		//fmt.Println(node.val)
		if val == node.val {
			//fmt.Println("repeat", node.val, val)
			return
		}

		if val < node.val {
			if node.left == nil {
				n := newNode(val)
				node.left = n
				n.parent = node
				return
			} else {
				node = node.left
			}
		} else {
			if node.right == nil {
				n := newNode(val)
				node.right = n
				n.parent = node
				return
			} else {
				node = node.right
			}
		}
	}
}

func (b *BiTree) Del(val int) {
	node := b.root
	for node != nil {
		if node.val == val {
			parent := node.parent
			// 如果该节点没有子孩子节点
			if node.left == nil && node.right == nil {
				if node == parent.left {
					parent.left = nil
				} else {
					parent.right = nil
				}
			}
			// 只有右孩子节点
			if node.left == nil && node.right != nil {
				if node == parent.left {
					parent.left = node.right
				} else {
					parent.right = node.right
				}
			}

			// 只有左孩子节点
			if node.left != nil && node.right == nil {
				if node == parent.left {
					parent.left = node.left
				} else {
					parent.right = node.left
				}
			}

			// 左右孩子都存在
			if node.left != nil && node.right != nil {

			}
		}
		if val < node.val {
			node = node.left
		} else {
			node = node.right
		}
	}
}

// 前序: 先访问根结点，再先序遍历左子树，最后再先序遍历右子树
func (n *node) printPreOrder() {
	if n != nil {
		fmt.Printf("%d,", n.val)
		n.left.printPreOrder()
		n.right.printPreOrder()
	}
}

// 后续: 先后序遍历左子树，然后再后序遍历右子树，最后再访问根结点
func (n *node) printPostOrder() {
	if n != nil {
		n.left.printPostOrder()
		n.right.printPostOrder()
		fmt.Printf("%d,", n.val)
	}
}

// 中序
func (n *node) printMidOrder() {
	if n != nil {
		n.left.printMidOrder()
		fmt.Printf("%d,", n.val)
		n.right.printMidOrder()
	}
}

func (b *BiTree) PrintPreOrder() {
	b.root.printPreOrder()
	println()
}

func (b *BiTree) PrintPostOrder() {
	b.root.printPostOrder()
	println()
}

func (b *BiTree) PrintMidOrder() {
	b.root.printMidOrder()
	println()
}
