package basicdsa

import "fmt"

//input : {1,2,3,4,5,6,7}, k:= 3
//output : {5,6,7,1,2,3,4}

func ReverseAndRotate(arr []int, k int) {
	n := len(arr)
	Reverse(arr, 0, n-k-1)
	Reverse(arr, n-k, n-1)
	Reverse(arr, 0, n-1)
	fmt.Println(arr)
}

func Reverse(arr []int, start, end int) {

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
