package main

import (
	"fmt"
	"go-code/binarytree"
)

func main() {
	// 创建测试二叉树:
	//     1
	//    / \
	//   2   3
	root := &binarytree.TreeNode{
		Val: 1,
		Left: &binarytree.TreeNode{Val: 2},
		Right: &binarytree.TreeNode{Val: 3},
	}

	// 调用中序遍历函数
	result := binarytree.InorderTraversal(root)

	// 打印结果
	fmt.Printf("中序遍历结果: %v\n", result)
	fmt.Printf("结果长度: %d\n", len(result))
}
