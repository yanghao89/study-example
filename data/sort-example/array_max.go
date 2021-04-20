package sort_example

import "math"

type ArrayMax struct {

}

func NewArrayMax()*ArrayMax  {
	return &ArrayMax{}
}

func (a *ArrayMax) GetMax(arr []int) int  {
	return a.process(arr,0,len(arr)-1)
}
// T(n) = aT(N/b) + O(N ^d)  a b d 都是常数   ----->  T(n) = 2T(N/2) + O(n ^ 0)

// T(n) = a * T(N/b) + O(N ^ d)


// log 以 b 为底 a次方 > d  --> O(N ^ log以b底 a次方)

// log 以 b 为底 a次方 < d ---> O(N^d)

// log 以 b 为底 a 次方 == d ----> O(N ^d * log以n为底)
//Max 获取数组最大值 时间复杂度 O(n)
func (a *ArrayMax) process(arr []int,L ,R int) int   {
	if L == R {
		return  arr[L]
	}
	min := L +((R - L ) >> 1)
	leftMax := a.process(arr,L,min) // N/2
	rightMax := a.process(arr,min+1,R) // N/2
	return int(math.Min(float64(leftMax),float64(rightMax)))
}