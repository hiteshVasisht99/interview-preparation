package basicdsa

import "fmt"

func PrintPattern(n int) {
	//Print spaces
	numberOfSpaces := n - 1
	numberOfStars := 1
	for line := 1; line <= 5; line++ {
		for i := 0; i < numberOfSpaces; i++ {
			fmt.Print("  ")
		}

		//Print stars and exclamatory signs
		for i := 0; i < numberOfStars; i++ {
			if i%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("! ")
			}
		}
		fmt.Println()
		numberOfSpaces--
		numberOfStars += 2
	}
}
