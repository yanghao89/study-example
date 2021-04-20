package sort_example

type TreeNode struct {
	Value int
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
	root := &TreeNode{Value: arr[min],Left: nil,Right: nil}
	root.Left = treeProcess(arr,L, min-1)
	root.Right = treeProcess(arr,min+1,R)
	return  root
}
func RangeSumBST(root *TreeNode,low,high int) int  {
	var (
		val int
	)
	bST(root,low,high,&val)
	return val
}
func bST(root *TreeNode,low,high int, num *int){
	if root == nil{
		return
	}
	bST(root.Left,low,high,num)
	if root.Value >= low && root.Value <= high{
		*num += root.Value
	}
	bST(root.Right, low,high,num)
}