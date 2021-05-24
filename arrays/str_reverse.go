package arrays

import (
	"strings"
)

// functin to reverse a string
// input: hello
// output: olleh
func Reverse(s string) string {
	// split s into an array of characters
	charactersArr := strings.Split(s, "")

	// loop through the array in the reverse order
	// and push each item to the new array
	strLen := len(charactersArr)
	strMaxIndex := strLen - 1
	reverseArr := make([]string, strLen)
	for i := strMaxIndex; i >= 0; i-- {
		reverseArr[strMaxIndex-i] = charactersArr[i]
	}

	// join all elements in the second array into a string
	return strings.Join(reverseArr, "")
}
