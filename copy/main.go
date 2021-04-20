package main

import (
	"fmt"
	"time"
)

type Robot struct {
	Name string `json:"name"`
	Age string `json:"age"`
	Count string `json:"count"`
}
func main()  {
	//bbb := make(map[int]int)
	//bbb[1]= 1
	//bbb[2] = 3

	fmt.Println(time.Now().UTC().Unix())



	//fmt.Printf("%d \n %d",[]int{12,3,4},bbb)
	/**
	深拷贝
	 */
	fmt.Printf("%s","深拷贝 内容是一样，改变其中一个对象的值时,另外一个不会发生变化。\n")
	robot1 := Robot{
		Name: "张三",
		Age: "12",
		Count: "测试1",
	}
	fmt.Printf("Robot1:%s\t内存地址:%p \n",robot1,&robot1)
	robot2 := robot1
	fmt.Printf("Robot2:%s\t内存地址:%p \n",robot2,&robot2)
	robot1.Name = "李四"
	fmt.Printf("Robot1:%s\t内存地址:%p \n",robot1,&robot1)
	fmt.Printf("Robot2:%s\t内存地址:%p \n",robot2,&robot2)

	/**
	浅拷贝
	 */
	fmt.Printf("%s","浅拷贝 内容和内存地址一样，改变其中一个对象的值时，另一个同时变化。\n")
	robot3 := Robot{
		Name: "张⑤",
		Age: "14",
		Count: "测试2",
	}
	robot4 := &robot3
	fmt.Printf("Robot3:%s\t内存地址:%p \n",robot3,&robot3)
	fmt.Printf("Robot4:%s\t内存地址:%p \n",robot4,robot4)
	robot3.Name = "张⑥"
	fmt.Printf("Robot3:%s\t内存地址:%p \n",robot3,&robot3)
	fmt.Printf("Robot4:%s\t内存地址:%p \n",robot4,robot4)

	/**

	 */
	fmt.Printf("%s","浅拷贝 使用new方式。\n")

	robot5 := new(Robot)
	robot5.Name = "张⑦"
	robot5.Age = "17"
	robot5.Count = "测试3"
	robot6 := robot5
	fmt.Printf("Robot5:%s\t内存地址:%p \n",robot5,robot5)
	fmt.Printf("Robot6:%s\t内存地址:%p \n",robot6,robot6)
	robot5.Name = "张⑧"
	fmt.Printf("Robot5:%s\t内存地址:%p \n",robot5,robot5)
	fmt.Printf("Robot6:%s\t内存地址:%p \n",robot6,robot6)





}
