package binarytree

func IsSymmetric(root *TreeNode) bool{
	// 1. 递归
/* 	if root==nil{
		return true
	}
		
	var f func(left *TreeNode,right *TreeNode) bool
	f = func(left, right *TreeNode) bool {
			if left==nil||right==nil{
				if left==nil&&right==nil {
					return true
				}else {
					return false
				}
			}

			if left.Val!=right.Val{return false}
			return  f(left.Right,right.Left)&&f(left.Left,right.Right)
	}

	return f(root.Left,root.Right) */

	// 2. 迭代
	var queue []int
	
}

