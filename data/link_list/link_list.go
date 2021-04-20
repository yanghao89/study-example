package link_list

import (
	"study-example/data/stack"
)

// IsPalindrome 判断是否为回文链表 1->2->3->2->1
func IsPalindrome( node *LinkedNode)bool  {
	if node == nil || node.Next == nil{
		return true
	}
	sk := stack.NewStack()
	tmp := node
	for tmp != nil{
		sk.Push(tmp.Value)
		tmp = tmp.Next
	}
	res := node
	for (!sk.IsEmpty()) && res != nil {
		if sk.Pop().(int) == res.Value{
			res = res.Next
		}else {
			return false
		}
	}
	return true
}

// 输入链表头节点,奇数长度返回中点,偶数长度返回上中点
func MinOrUpMidNode(node *LinkedNode) *LinkedNode  {
	if node == nil || node.Next == nil || node.Next.Next == nil{
		return node
	}
	l := node.Next
	r := node.Next.Next
	for r.Next != nil && r.Next.Next != nil{
		l = l.Next
		r = r.Next.Next
	}
	return l
}

// 输入链表头节点,奇数长度返回中点,偶数长度返回下中点
func MinOrDownMidNode(node *LinkedNode) *LinkedNode  {
	if node == nil || node.Next == nil{
		return node
	}
	l := node.Next
	r := node.Next
	for r.Next != nil && r.Next.Next != nil{
		l = l.Next
		r = r.Next.Next
	}
	return l
}


// 输入链表头节点,奇数长度返回中点前一个,偶数长度返回上中点前一个
func MinOrUpPreMidNode(node *LinkedNode) *LinkedNode {
	if node == nil || node.Next == nil || node.Next.Next == nil{
		return node
	}
	l := node
	r := node.Next.Next
	for r.Next != nil && r.Next.Next != nil{
		l = l.Next
		r = r.Next.Next
	}
	return l
}

// 输入链表头节点,奇数长度返回中点前一个,偶数长度返回下中点前一个
func MinOrUpDownMidNode(node *LinkedNode)*LinkedNode  {
	if node== nil || node.Next == nil{
		return node
	}
	if node.Next.Next == nil{
		return node
	}
	l := node
	r := node.Next
	for r.Next != nil && r.Next.Next != nil{
		l = l.Next
		r = r.Next.Next
	}
	return l
}

