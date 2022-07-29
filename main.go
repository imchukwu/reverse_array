package main

import (
	"fmt"
)

func main() {
	myArray := []int{1, 4, 5, 2, 7, 8, 9}
	left, right := 0, len(myArray)-1

	for i := 0; i < len(myArray)-1; i++ {
		if left < right {
			temp := myArray[left]
			myArray[left] = myArray[right]
			myArray[right] = temp
			left++
			right--
		} else {
			break
		}
	}
	fmt.Printf("Array Reverse: %v", myArray)

}
