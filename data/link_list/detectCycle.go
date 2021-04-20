package link_list

import "unsafe"


type ListNode struct {
	Val int
	Next *ListNode
}


//detectCycle  环形链表检测 https://leetcode-cn.com/problems/linked-list-cycle-lcci/
func detectCycle(head *ListNode) *ListNode {
	m := make(map[uintptr]*ListNode)
	for head != nil{
		if _,ok := m[uintptr(unsafe.Pointer(head))];ok{
			return head
		}
		m[uintptr(unsafe.Pointer(head))] = head
		head = head.Next
	}
	return nil
}
