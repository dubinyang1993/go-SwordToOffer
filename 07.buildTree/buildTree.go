package main

import "fmt"

// 重建二叉树
func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)
	fmt.Printf("tree: %+v\n", tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 根节点
	root := &TreeNode{preorder[0], nil, nil}
	// 核心：找到根节点在中序遍历中的位置
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 左子树
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	// 右子树
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
