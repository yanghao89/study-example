package sort_example

//插入排序 O(n ^ 2)
func InsertSort(arr []int)[]int  {
	for i := 1; i < len(arr);i++{
		for j:= i-1;j >= 0 && arr[j] > arr[j+1];j --{
			Swap(arr,j,j+1)
		}
	}
	return arr
}
