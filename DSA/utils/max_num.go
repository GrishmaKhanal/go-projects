package utils

import "fmt"

func LargestNumber() {
	numbers := []int{1, 2, 3, 4, 5}

	max_num := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if max_num < numbers[i] {
			max_num = numbers[i]
		}
	}
	fmt.Println("Max Num = ", max_num)
}
