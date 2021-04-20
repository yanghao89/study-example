package tree

import (
	"fmt"
	"study-example/data/stack"
)

type PrintTree struct {
}

func NewPrintTree() *PrintTree  {
	return &PrintTree{
	}
}

//先序遍历 头左右
func (n *PrintTree) Pre(node *TreeNode)  {
	if node == nil{
		return
	}
	fmt.Printf("%d ",node.Val)
	n.Pre(node.Left)
	n.Pre(node.Right)
}

//中序遍历 左头右
func (n *PrintTree) In(node *TreeNode)  {
	if node == nil{
		return
	}
	n.In(node.Left)
	fmt.Printf("%d ",node.Val)
	n.In(node.Right)
}

//后序遍历   左右头
func (n * PrintTree) Hn(node *TreeNode)  {
	if node == nil{
		return
	}
	n.Hn(node.Left)
	n.Hn(node.Right)
	fmt.Printf("%d ",node.Val)
}

//中序遍历 左头右
func (n *PrintTree) IterIn(node *TreeNode)  {
	if node == nil{
		return
	}
	sk :=  stack.NewStack()
	for (!sk.IsEmpty()) || (node != nil){
			if node != nil{
				sk.Push(node)
				node = node.Left
			}else {
				node = sk.Pop().(*TreeNode)
				fmt.Printf("%d ",node.Val)
				node = node.Right
			}
	}
}

//先序遍历 迭代 头左右
func  (n *PrintTree) IterPre(node *TreeNode)  {
	sk := stack.NewStack()
	sk.Push(node)
	for !sk.IsEmpty(){
			node = sk.Pop().(*TreeNode)
			fmt.Printf("%d ",node.Val)
			if node.Right != nil{
				sk.Push(node.Right)
			}
			if node.Left != nil{
				sk.Push(node.Left)
			}
	}
}

func  (n *PrintTree) IterHn(node *TreeNode)  {
	sk := stack.NewStack()
	sk.Push(node)
	sk1 := stack.NewStack()
	for !sk.IsEmpty(){
		node = sk.Pop().(*TreeNode)
		sk1.Push(node)
		if node.Left != nil{
			sk.Push(node.Left)
		}
		if node.Right != nil{
			sk.Push(node.Right)
		}
	}
	for !sk1.IsEmpty(){
		node2 := sk1.Pop().(*TreeNode)
		fmt.Printf("%d ",node2.Val)
	}
}