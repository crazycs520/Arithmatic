package btree

import (
	"container/list"
	"dsa_thu/stack"
)

type BinTreeNode struct {
	data      interface{}
	leftNode  *BinTreeNode
	rightNode *BinTreeNode
}

//先序遍历, 递归实现
func TravPre_0(root *BinTreeNode, visit func(node *BinTreeNode)) {
	if root == nil {
		return
	}
	visit(root)
	TravPre_0(root.leftNode, visit)
	TravPre_0(root.rightNode, visit)
}

//先序遍历, 迭代实现
func TravPre_1(x *BinTreeNode, visit func(node *BinTreeNode)) {
	var s *stack.Stack // 辅助栈
	for {
		visitAlongLeftBranch(x, visit, s)
		if s.IsEmpty() {
			break
		}
		a, _ := s.Pop()
		x = a.(*BinTreeNode)
	}

}

func visitAlongLeftBranch(x *BinTreeNode, visit func(node *BinTreeNode), s *stack.Stack) {
	for x != nil {
		visit(x)
		s.Push(x.rightNode)
		x = x.leftNode
	}
}

//中序遍历 递归实现
func TravIn_0(x *BinTreeNode, visit func(node *BinTreeNode)) {
	if x == nil {
		return
	}
	TravIn_0(x.leftNode, visit)
	visit(x)
	TravIn_0(x.rightNode, visit)
}

//中序遍历 迭代实现
func TravIn_0(x *BinTreeNode, visit func(node *BinTreeNode)) {
	var s *stack.Stack
	for {
		goAlongLeftBranch(x, s)
		if s.IsEmpty() {
			break
		}
		a, _ := s.Pop()
		x = a.(*BinTreeNode)
		visit(x)
		x = x.rightNode
	}
}

func goAlongLeftBranch(x *BinTreeNode, s *stack.Stack) {
	for x != nil {
		s.Push(x)
		x = x.leftNode
	}
}

//层次遍历
func TravLevel(x *BinTreeNode, visit func(node *BinTreeNode)) {
	l := list.New()
	l.PushBack(x) //根节点入队
	for l.Len() != 0 {
		a := l.Front()
		x = a.Value.(*BinTreeNode)
		visit(x)
		if x.leftNode != nil {
			l.PushBack(x.leftNode) //坐孩子入队
		}
		if x.rightNode != nil {
			l.PushBack(x.rightNode)
		}
	}

}

//普通二叉树求深度
func treeDeep(root *BinTreeNode) int {
	if root == nil {
		return 0
	}
	nleft := treeDeep(root.leftNode)
	nright := treeDeep(root.rightNode)
	if nleft > nright {
		return nleft + 1
	}
	return nright + 1
}

//判断是否是二差平衡树--低效
func isBalanced(root *BinTreeNode) bool {
	if root == nil {
		return true
	}
	left := treeDeep(root.leftNode)
	right := treeDeep(root.rightNode)
	diff := left - right
	if diff > 1 || diff < -1 {
		return false
	}
	return isBalanced(root.leftNode) && isBalanced(root.rightNode)
}

//判断是否是二差平衡树--高效
func isBalanced2(root *BinTreeNode) bool {
	depth := 0
	return isBalanced2_(root, &depth)
}

func isBalanced2_(root *BinTreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}
	left, right := 0, 0
	if isBalanced2_(root.leftNode, &left) && isBalanced2_(root.rightNode, &right) {
		diff := left - right
		if diff <= 1 && diff >= -1 {
			if left > right {
				*depth = left + 1
			} else {
				*depth = right + 1
			}
			return true
		}
	}

	return false
}
