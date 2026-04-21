package main

import (
	"fmt"
	basicdsa "interview/basic-dsa"
)

func main() {
	fmt.Println(basicdsa.CanCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}
