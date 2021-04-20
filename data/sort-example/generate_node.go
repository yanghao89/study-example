package sort_example

type LinkedNode struct {
	Value int `json:"value"`
	Next *LinkedNode `json:"next"`
}

func GenerateNode(arr []int) *LinkedNode {
	if len(arr) <= 1{
		return nil
	}
	res := &LinkedNode{
		Value: arr[0],
	}
	tmp := res
	for i:=1;i < len(arr);i++{
		node := &LinkedNode{Value: arr[i]}
		tmp.Next = node
		tmp = node
	}
	return res
}

type DoubleNode struct {
	Value int `json:"value"`
	Last *DoubleNode `json:"last"`
	Next *DoubleNode `json:"next"`
}

func GenerateNodeList(arr []int) *DoubleNode  {
	if len(arr) <= 1{
		return nil
	}
	res := &DoubleNode{
		Value: arr[0],
		Last: nil,
	}
	tmp := res
	for i:=1; i< len(arr);i++{
		node := &DoubleNode{
			Value: arr[i],
			Last: tmp,
		}
		tmp.Next = node
		tmp = node
	}
	return res
 }
  