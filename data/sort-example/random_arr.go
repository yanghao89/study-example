package sort_example

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(maxSize,maxValue int)(arr []int)  {
	rand.Seed(time.Now().Unix())
	arr =  make([]int,rand.Intn(maxSize))
	for i:=0;i < len(arr);i++{
		arr[i] = ((maxValue +1) * rand.Intn(maxSize)) - (maxValue * rand.Intn(maxSize))
	}
	return  arr
}
