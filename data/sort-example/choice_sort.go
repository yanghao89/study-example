package sort_example

//ChoiceSort 选择排序 时间复杂度为 O(n^2)
func ChoiceSort (arr []int)[]int  {
	for i:=0;i < len(arr)-1; i++{
		min := i
		for j:=i+1; j < len(arr);j++{
			if arr[j] < arr[min] {
				min = j
			}
		}
		Swap(arr,i,min)
	}
	return arr
}
