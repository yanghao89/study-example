package recursion

import "study-example/data/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *tree.TreeNode {
	return bst(nums,0,len(nums)-1)
}

func bst(nuns []int,L,R int) *tree.TreeNode  {
	if L > R {
		return nil
	}
	min := L +((R - L) >> 1)
	root := &tree.TreeNode{Val: nuns[min]}
	root.Left = bst(nuns,L,min -1)
	root.Right = bst(nuns,min+1,R)
	return root
}