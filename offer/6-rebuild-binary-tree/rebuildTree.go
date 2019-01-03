package rebuildTree

import "fmt"

// 给定前序遍历，如 { 1,2,4,7,3,5,6,8}
//    中序遍历，如  { 4,7,2,1,5,3,8,6}
//重建该二叉树，并返回更节点

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func BuildTreeBySlice(list []int) *BinaryTree {
	if list == nil || len(list) == 0{
		return nil
	}
	root := &BinaryTree{}
	i:=0
	node := root
	for {
		node.value=list[i]
		left := &BinaryTree{
			value:list
		}

	}
}

func PrintLevelBinaryTree(root *BinaryTree) {
	list := make([]*BinaryTree , 0)
	if root == nil {
		return
	}
	list = append(list,root)
	num := 1
	for len(list) > 0 {
		node := list[0]
		list = list[1:]
		num--
		fmt.Print("%d ",node.value)
		if num == 0{
			fmt.Print("\n")
		}
		if node.left != nil {
			list = append(list,node.left)
		}
		if node.right != nil {
			list = append(list,node.right)
		}
	}
}

func RebuildBinaryTree(preOrder , inOrder []int ) * BinaryTree {
	if preOrder == nil || inOrder == nil || len(preOrder) != len(inOrder){
		return nil
	}
	return constructCore(preOrder,inOrder,0,len(preOrder) , 0 , len(inOrder))
}

func constructCore( preOrder , inOrder []int , startPre,endPre int, startIn , endIn int) *BinaryTree{
	//前序遍历的第一个数值是根节点的值
	rootValue := preOrder[0]
	root := &BinaryTree{}
	root.value = rootValue

	if startPre == endPre {
		if startIn == endIn && preOrder[startPre] == inOrder[startIn] {
			return root
		}else{
			// error invalid input
		}
	}
	i := startIn
	for ; i < endIn ; i++ {
		if inOrder[i] == rootValue {
			break
		}
	}
	if i == endIn && inOrder[i] != rootValue {
		// error invalid input
	}
	leftLength := i - startIn
	if leftLength > 0{
		//构建左子树
		root.left = constructCore(preOrder , inOrder , startPre +1, startPre + leftLength , startIn , i - 1)
	}
	if leftLength < (endPre - startPre) {
		//构建右子树
		root.right = constructCore(preOrder , inOrder , startPre + leftLength +1, endPre , i+1 , endIn)
	}
	return root
}

