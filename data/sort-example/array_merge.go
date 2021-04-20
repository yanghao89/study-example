package sort_example

import (
	"fmt"
	"sort"
)

func MergeTwoSorted(s1 []int,s2 []int) {
	s4:= s2
	s3 := append(s1,s2...)
	s3 = append(s3,s4...)
	sort.Ints(s3)
	fmt.Println(s3)
}
