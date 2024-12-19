package utils

import "fmt"

func Sum(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	fmt.Printf("The Sum of %v natural numbers is: %v\n", n, sum)
}
