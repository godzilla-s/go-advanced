package tree

// 总结平衡二叉树特点：
//1. 非叶子节点最多拥有两个子节点；
//2. 非叶子节值大于左边子节点、小于右边子节点；
//3. 树的左右两边的层级数相差不会大于1;
//4. 没有值相等重复的节点;

type treeNode struct {
	hash  int
	val   interface{}
	left  *treeNode
	right *treeNode
}

type AVLTree struct {
	root *treeNode
}

func newTreeNode(hash int, val interface{}) *treeNode {
	return &treeNode{
		hash: hash,
		val:  val,
	}
}

func (t *AVLTree) leftRotate() {

}

func (t *AVLTree) rightRotate() {

}

func (t *AVLTree) Add(hash int) {
	if t.root == nil {
		t.root = newTreeNode(hash, nil)
		return
	}

	root := t.root
	if hash < root.hash {

	} else if hash > root.hash {

	} else {

	}
}
