package binarytree


// InorderTraversal 给定一个二叉树的根节点 root ，返回它的中序遍历 。
func InorderTraversal(root *TreeNode) []int {
	// 1. 递归方法，使用系统栈
/* 	var res []int
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root==nil {
		return
	}
	traversal(root.Left)
	res =  append(res,root.Val)
	traversal(root.Right)
	}
	traversal(root)
	return res */

	// 2. 迭代方法，模拟栈
	
	var res []int
      var stack []*TreeNode
      cur := root
	var lastVisited *TreeNode
      for cur != nil || len(stack) > 0 {
          // 1. 一路向左，压栈
          for cur != nil {
			stack = append(stack, cur)	// 压栈,用于保存回溯的节点
            cur = cur.Left
          }

          // 2. 弹栈
		  node := stack[len(stack)-1]

		  if node.Right==nil || node.Right==lastVisited{
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]// update stack
			lastVisited = node
		  }else{
			// 3. 转向右子树
          	cur = node.Right
		  }
      }

      return res
}