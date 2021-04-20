package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	str = `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgs+j/FitdkUrRS4tQ
Xcs7FTq/LARkVPH99sxuE5A1UIWgCgYIKoZIzj0DAQehRANCAARgwgz11yuw/8Hi
Qd4NuKzsw4ZYp3Q6v1xhnprQn8XA0yoKM1AsUqtlEoGeubqepUxHgloyTG4XYjuq
rdYRSRR8
-----END PRIVATE KEY-----`
)

func connHandler(c net.Conn)  {
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			return
		}

		c.Write([]byte(input))

		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			continue
		}

		fmt.Print(string(buf[0:cnt]))
	}
}
func main() {
	conn, err := net.Dial("tcp", "52.74.223.119:80")
	if err != nil {
		fmt.Printf("Fail to connect, %s\n", err)
		return
	}

	connHandler(conn)

	//
	//a := fmt.Sprintf("%.2f%%", 123.456)
	//fmt.Println(a)
	//arr0 := sortExample.ChoiceSort([]int{77,33,44,25,88,99,95,23,56})
	//fmt.Println(arr0)
	//
	//arr1 := sortExample.BubbleSort([]int{77,33,44,25,88,99,95,23,56})
	//arr2 := sortExample.InsertSort(sortExample.GenerateRandomArray(100,100))
	//fmt.Println(arr2)

	//fmt.Println(a)
	//fmt.Printf("%b",a)
	//
	//fmt.Printf("%b", 1020203%10)
	//for i := 1;i<1000;i++{
	//	    fmt.Println(i,"\n")
	//
	//}
	//a - (a / b) * b
	//node := sort_example.GenerateNode([]int{77,33,44,25,88,99,95,23,56})
	///fmt.Println(sort_example.GenerateNode([]int{77,33,44,25,88,99,95,23,56}))
	//sort_example.FormtNode(node)
	//node = sort_example.ReverseNode(node)

	//nodes := sort_example.GenerateNode([]int{77,33,44,25,33,88,99,95,23,56,33,33,33,33})
	//fmt.Printf("%v\n",nodes)
	////nodes = sort_example.ReverseNodeList(nodes)
	////fmt.Printf("%v\n",node)
	//fmt.Printf("%v\n",nodes)
	//
	//fmt.Println("%v\n", sort_example.RemoveNodeList(nodes,33))
	//fmt.Printf("%x\n",891)
	//fmt.Printf("%x\n",10-1)
	//	stack := sort_example.NewStack()
	//queue := sort_example.NewQueue()
	//	for _,v := range []int{77,33,44,25,88,99,95,23,56}{
	//		stack.Push(v)
	//		queue.Push(v)
	//	}
	//	//for _,v := range []int{77,33,44,25,88,99,95,23,56}{
	//	//	queue.AddFromTail(v)
	//	//}
	//	fmt.Println(stack.Pop(),stack.Count(),queue.Poll(),stack.Count())

	//array := sort_example.NewArrayMax()
	//max := tree.ArrayToCreateTree([]int{77,33,44,25,88,99,95,56,23,24})
	//
	//printTree := tree.NewPrintTree()
	//fmt.Printf("%s ","中序递归---->左头右")
	//printTree.In(max)
	//fmt.Printf("\n%s ","中序迭代---->左头右")
	//printTree.IterIn(max)
	//fmt.Printf("\n%s ","前序递归----->头左右")
	//printTree.Pre(max)
	//fmt.Printf("\n%s ","前序迭代----->头左右")
	//printTree.IterPre(max)
	//fmt.Printf("\n%s ","后序递归----->左右头")
	//printTree.Hn(max)
	//fmt.Printf("\n%s ","后序迭代----->左右头")
	//printTree.IterHn(max)

	//authKey, err := token.AuthKeyFromBytes([]byte(str))
	//if err != nil {
	//	log.Fatal("token error:", err)
	//}
	////fmt.Printf("%s", authKey)
	//
	//for _, v := range []string{"e50d11d823d6180ccbde9d063073c8fbc0fa8966e0e9e12d7026287da73b2224", "709657e91ea26fd00b398657ebc730b9e72b2c693776915cb458bca3fe71bc9e", "eb72ca74795ce23386f8509428e265e751a8009bc8679a4c7d40545b6092fe0d",
	//	"709657e91ea26fd00b398657ebc730b9e72b2c693776915cb458bca3fe71bc9e", "117db7a7a569de28bdff0059544c1dfe376e6df66ff6d6750af6c30ed9306309", "b4cbf1447624d9345bc1586982782fdd295cf050614165e28dae703835d813f6", "34b97df0073f8731aa35564c8d166d7e396b3f23504e08b6ff6f97415686d587",
	//	"eb72ca74795ce23386f8509428e265e751a8009bc8679a4c7d40545b6092fe0d", "b4cbf1447624d9345bc1586982782fdd295cf050614165e28dae703835d813f6"}{
	//notification := &apns2.Notification{}
	//notification.DeviceToken = v
	//notification.Topic = "com.luckludoapps.ios"
	//notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`)
	//tokens := &token.Token{
	//	AuthKey: authKey,
	//	// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
	//	KeyID: "MKNK8JY4A8",
	//	// TeamID from developer account (View Account -> Membership)
	//	TeamID: "3D6J3XY8L8",
	//}
	//	client := apns2.NewTokenClient(tokens)
	//	res, err := client.Push(notification)
	//
	//	if err != nil {
	//		log.Fatal("Error:", err)
	//	}
	//
	//	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
	//
	//}

	//node := link_list.CreateLinkedNode([]int{1,2,3,2,1})
	//node := link_list.CreateLinkedNode([]int{1,2,3,4,5,6,7,8,9})
	//is1 :=link_list.MinOrUpMidNode(node)
	//is2 := link_list.MinOrDownMidNode(node)
	//is3 := link_list.MinOrUpPreMidNode(node)
	//is4 := link_list.MinOrUpDownMidNode(node)
	//i5 := link_list.CreateLinkNode([]int{77,33,44,25,88,99,95,56,23,24})
	//fmt.Println(i5)
	//bb := link_list.DeepCopyLinkedList(i5)
	//fmt.Println(bb)
	//fmt.Println(i5)

			//fmt.Println(4096 * 2/ 1024)




	//arr :=[8]int{}
	//for i := 0; i < 8; i++ {
	//	arr[i] = i
	//	//arr = append(arr,i)
	//}
	//
	//fmt.Println(arr)
	//exchange(&arr)
	//fmt.Println(arr)
}


func exchange(arr *[8]int) {
	for k, v := range arr {

		arr[k] = v * 2
	}
}