package sort_example

func BinarySearch(arr []int,key int)int  {

	index,l,r := -1 ,0,len(arr) - 1
	for l < r{
		mid := l + ((r -l)>>1)
		if arr[mid] >= key {
			index = mid
			r = mid -1
		}else {
			l = mid +1
		}
	}
	return index
}
