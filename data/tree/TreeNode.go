package tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func ArrayToCreateTree(arr []int) *TreeNode   {
	return treeProcess(arr, 0,len(arr)-1)
}

func treeProcess(arr []int, L,R int) *TreeNode  {
	if L > R {
		return  nil
	}
	min := L +((R - L ) >> 1)
	root := &TreeNode{Val: arr[min],Left: nil,Right: nil}
	root.Left = treeProcess(arr,L, min-1)
	root.Right = treeProcess(arr,min+1,R)
	return  root
}