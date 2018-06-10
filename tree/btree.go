// B树

// 特征：
// 根节点至少有两个子节点
// 每个节点有M-1个key，并且以升序排列
// 位于M-1和M key的子节点的值位于M-1 和M key对应的Value之间
// 其它节点至少有M/2个子节点

package tree

type btree_node struct {
	num    int
	keys   []int
	child  []*btree_node
	parent *btree_node
}

type BTree struct {
	order int
	root  *btree_node
}

func NewBtree(order int) *BTree {
	return &BTree{
		order: order,
	}
}

func (bt *BTree) Add(key int) {
	if bt.root == nil {
		node := new(btree_node)
		node.keys = make([]int, bt.order)
		node.child = make([]*btree_node, bt.order+1)

		node.num = 1
		node.keys[0] = key
		bt.root = node
		return
	}

	root := bt.root
	idx := 0
	for root != nil {
		for idx = 0; idx < root.num; idx++ {
			if key < root.keys[idx] {
				break
			}
		}

		if root.child[idx] != nil {
			root = root.child[idx]
		} else {
			break
		}
	}

	bt.add(root, key, idx)
}

func (bt *BTree) add(node *btree_node, key, idx int) {
	for i := node.num; i > idx; i-- {
		node.keys[i] = node.keys[i-1]
	}

	node.keys[idx] = key
}
