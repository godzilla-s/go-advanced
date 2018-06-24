// B 树
package tree

import "fmt"

type Btree struct {
	m      int      // 阶数
	num    int      // 个数 : max(num) = m - 1
	key    []int    // 值
	parent *Btree   // 指向父亲的节点
	child  []*Btree // 子节点
}

func New(m int) *Btree {
	return &Btree{
		m:     m,
		key:   make([]int, m), // 留一个做备用
		child: make([]*Btree, m),
	}
}

func (b *Btree) Add(val int) {
	n := b
	i := 0
	//parent := b.parent
	for {
		for i = 0; i < n.num; i++ {
			if val < b.key[i] {
				//copy(b.key[i+1:], b.key[i:])
				// n = n.child[i]
				break
			}
		}
		if n.child[i] == nil {
			break
		}
		n = n.child[i]
	}

	copy(n.key[i+1:], n.key[i:])
	n.key[i] = val
	fmt.Println(n.key)
	n.num++
	if n.num == n.m {
		// 分裂
		fmt.Println("split")
		idx := n.num / 2
		b.split(idx)
	}
}

// 分裂
func (b *Btree) split(idx int) {
	p := New(b.m)
	copy(p.key, b.key[idx:])
	p.num = len(p.key)
	if b.parent == nil {
		node := New(b.m)
		node.key[0] = b.key[idx]

	}
}
