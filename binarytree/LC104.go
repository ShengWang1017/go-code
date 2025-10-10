package binarytree

import "math"

func MaxDepth(root *TreeNode) int{
	if root ==nil {
		return 0
	}

	leftMax := MaxDepth(root.Left)
	rightMax := MaxDepth(root.Right)
	max := int(math.Max(float64(leftMax),float64(rightMax)))
	max++
	return max
}