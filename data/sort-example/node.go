package sort_example

//ReverseNode 单向链表反转
func ReverseNode(head *LinkedNode) *LinkedNode  {
	pre,next:= &LinkedNode{},&LinkedNode{}
	for head != nil{
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

//ReverseNodeList  双向链表反转
func ReverseNodeList(head *DoubleNode) *DoubleNode  {
	pre,next:= &DoubleNode{}, &DoubleNode{}
	for head != nil{
		next = head.Next
		head.Next = pre
		head.Last = next
		pre = head
		head  = next
	}
	return pre
}

//RemoveNodeList 删除指定链表中的值
func RemoveNodeList(head *LinkedNode,num int) *LinkedNode  {
	for head != nil{
		if head.Value != num{
			break
		}
		head = head.Next
	}
	pre,cur:= head,head
	for cur != nil{
			if cur.Value == num{
				pre.Next = cur.Next
			}else {
				pre = cur
			}
			cur = cur.Next
	}
	return head
}
