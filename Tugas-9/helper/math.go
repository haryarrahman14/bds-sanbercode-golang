package helper

import (
	"fmt"
	"strings"
)

func Sum(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func FormatNumbers(numbers ...[]int) string {
	var formattedNumbers []string
	for _, nums := range numbers {
		formattedNumbers = append(formattedNumbers, FormatSlice(nums)...)
	}
	return strings.Join(formattedNumbers, " + ")
}

func FormatSlice(numbers []int) []string {
	var formattedNumbers []string
	for _, num := range numbers {
		formattedNumbers = append(formattedNumbers, fmt.Sprint(num))
	}
	return formattedNumbers
}
