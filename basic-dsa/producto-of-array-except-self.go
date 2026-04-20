package basicdsa

func ProductOfArrayExceptSelf(arr []int) []int {
	left := make([]int, len(arr))
	right := make([]int, len(arr))

	left[0] = 1
	for i := 1; i < len(arr); i++ {
		left[i] = left[i-1] * arr[i-1]
	}

	right[len(right)-1] = 1
	for i := len(arr) - 2; i >= 0; i-- {
		right[i] = right[i+1] * arr[i+1]
	}

	for i := 0; i < len(left); i++ {
		left[i] = left[i] * right[i]
	}
	return left

}
