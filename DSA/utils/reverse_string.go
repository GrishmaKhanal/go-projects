package utils

import "fmt"

func ReverseString(input_string string) {
	len := len(input_string)
	// var output_string string // appending to strings in Go creates a new string each time (due to immutability)
	output_string := make([]byte, len)

	for i := 0; i < len; i++ {
		output_string[len-i-1] += input_string[i]
	}
	fmt.Println(string(output_string))
}
