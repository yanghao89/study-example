package sort_example

// BubbleSort 时间复杂度O(n ^2)   空间复杂度 O(1)
func BubbleSort(arr []int)[]int  {
	for i:= len(arr)-1; i > 0 ;i--{
		for j:=0;j < i;j++{
			if arr[j] > arr[j+1]{
				Swap(arr,j,j+1)
			}
		}
	}
	return arr
}

func Swap(arr []int,i,j int)  {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
