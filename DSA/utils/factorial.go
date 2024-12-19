package utils

import "fmt"

func Factorial(n int) {
	factorial := 1
	if n < 2 {
		fmt.Println(factorial)
		return
	}
	for i := 1; i <= n; i++ {
		factorial = factorial * i
	}
	fmt.Println(factorial)
	return
}
