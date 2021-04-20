package link_list

type LinkedNode struct {
	Value int `json:"value"`
	Next *LinkedNode `json:"next"`
}

func CreateLinkedNode(arr []int)*LinkedNode  {
	node := &LinkedNode{Value: arr[0]}
	tmp := node
	for i:= 1; i<len(arr);i++{
		ne := &LinkedNode{Value: arr[i]}
		tmp.Next = ne
		tmp = ne
	}
	return node
}

type LinkNode struct {
	Value int `json:"value"`
	Next *LinkNode `json:"next"`
	Rand *LinkNode `json:"rand"`
}

func CreateLinkNode(arr []int)*LinkNode  {
	node := &LinkNode{Value: arr[0]}
	tmp := node
	for i:=1;i < len(arr);i++{
		ne := &LinkNode{Value: arr[i]}
		rand := ne
		for n:= 0; n < len(arr);n += 2{
				ra := &LinkNode{Value: arr[i]}
				rand.Rand = ra
				rand = ra
		}
		tmp.Next = ne
		tmp.Rand = rand
		tmp = ne
	}
	return node
}
