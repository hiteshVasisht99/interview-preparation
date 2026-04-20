package basicdsa

import "fmt"

func PrintFibonacci(n int) {
	a := 0
	b := 1

	for i := 1; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	fmt.Printf("%dth number of fibonacci is: %d\n", n, a)
}
