package main

import (
	"fmt"
	"sync"
)

//func main() {
//	size := 10
//	ch := make(chan []int)
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//	go printSlice(ch, wg)
//
//	list := make([]int, 0, size)
//
//	for i := 0; i < 100; i++ {
//		if len(list) == size {
//			ch <- list
//			list = list[:0]
//		}
//		list = append(list, i)
//	}
//	//if len(list) > 0 {
//	//	ch <- list
//	//}
//	close(ch)
//	wg.Wait()
//
//	fmt.Println("Hello, playground")
//}a

func main()  {
	//a := "345678911002233"
	//// 0 1000000
	//
	////00 110011
	//
	//for _,v := range a{
	//	fmt.Println("\n")
	//	fmt.Printf("%b",v)
	//}
	//t := time.Now().UTC().AddDate(0, 0, -1).Format("2006-01-02")
	//fmt.Println(t)
	////t1, _ := time.Parse("2006-01-02 15:04:05","2021年 03月 19日 07:16:02")
	//
	//
	//tm2, _ := time.Parse("2006-01-02", t)
	//
	//fmt.Println(tm2.Unix(),1616112000 + 86399)

	//for _,v := range []int{4,5,6,7}{
	//	var i int
	//	for _,vv:= range []float64{1.10,2.15,3.45,4.65}{
	//		i += int(vv * 100)
	//	}
	//
	//}
	//likes := fmt.Sprintf("like %s'%'", "ca")
	//fmt.Println(likes)
	//a := []rune("-27")
	//fmt.Println(a)



	//fmt.Println(t1.Unix(),time.Now().AddDate(0,0, -1).Unix())
}

func printSlice(ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		s, ok := <-ch
		if ok {
			fmt.Printf("received: %+v\n", s)
		} else {
			break
		}
	}
}

