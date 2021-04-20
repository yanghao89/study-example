package link_list

import "unsafe"


type Node struct {
  Val int
	Next *Node
  Random *Node
}

// CopyRandomList 复制带随机指针的链表 https://leetcode-cn.com/problems/copy-list-with-random-pointer/
func CopyRandomList(head *Node) *Node {
	tmp := head
	maps := make(map[uintptr]*Node)
	for tmp != nil{
		maps[uintptr(unsafe.Pointer(tmp))] = &Node{Val: tmp.Val}
		tmp = tmp.Next
	}
	tmp = head
	for tmp != nil{
		maps[uintptr(unsafe.Pointer(tmp))].Next = maps[uintptr(unsafe.Pointer(tmp.Next))]
		maps[uintptr(unsafe.Pointer(tmp))].Random =  maps[uintptr(unsafe.Pointer(tmp.Random))]
		tmp = tmp.Next
	}
	return maps[uintptr(unsafe.Pointer(head))]
}

//func CircularLinkedList(head1,head2 *LinkedNode) *LinkedNode {
//	isCircular := func(head *LinkedNode) *LinkedNode {
//		m := make(map[uintptr]*LinkedNode)
//		tmp := head
//		for tmp != nil {
//			_, ok := m[uintptr(unsafe.Pointer(tmp))]
//			if !ok{
//				m[uintptr(unsafe.Pointer(tmp))] = tmp
//			}else {
//				return tmp
//			}
//			tmp = tmp.Next
//		}
//		return nil
//	}
//	node1 := isCircular(head1)
//	if isCircular(head1) != nil{
//		return  node1
//	}
//	node1 = isCircular(head2)
//	if node1 != nil {
//		return node1
//	}
//	return nil
//}


