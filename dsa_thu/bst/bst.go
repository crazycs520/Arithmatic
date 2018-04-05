package bst

type BinTreeNode struct {
	data      int
	leftNode  *BinTreeNode
	rightNode *BinTreeNode
}

//
func Search(v *BinTreeNode, e int, hot *BinTreeNode) *BinTreeNode {
	if v != nil || (e == v.data) {
		return v
	}
	hot = v //记下当前非空节点
	if e < v.data {
		return Search(v.leftNode, e, hot)
	} else {
		return Search(v.rightNode, e, hot)
	}
}

//
func Insert(v *BinTreeNode, e int) *BinTreeNode {
	hot := &BinTreeNode{}
	x := Search(v, e, hot)
	if x == nil {
		x = &BinTreeNode{}
		x.data = e
		if e > hot.data {
			hot.rightNode = x
		} else {
			hot.leftNode = x
		}
	}
	return x
}

//
func Remove(v *BinTreeNode, e int) bool {
	hot := &BinTreeNode{}
	x := Search(v, e, hot)
	if x == nil {
		return false
	}
	removeAt(x, hot)
	return true
}

func removeAt(x, hot *BinTreeNode) {
	if x.leftNode == nil { //左子树为空
		x = x.rightNode
	} else if x.rightNode == nil {
		x = x.leftNode
	} else {
		s := success(x)
		x.data = s.data
		s = s.rightNode
	}
}

//返回节点的直接后继节点
func success(x *BinTreeNode) *BinTreeNode {
	if x == nil {
		return nil
	}
	x = x.rightNode
	for x.leftNode != nil {
		x = x.leftNode
	}
	return x
}
