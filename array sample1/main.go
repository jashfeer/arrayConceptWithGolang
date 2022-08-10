package main

import "fmt"

func sample(arr [9]int, match int) (int, int) {
	for i := 0; i < (len(arr) - 1); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == match {
				return arr[i], arr[j]
			}
		}
	}
	return 0, 0
	//O(1)S
	//O(n^2)T
}

func main() {
	var number1, number2 int
	match := 10
	array := [...]int{2, 6, 4, 7, 9, 2, 5, 7,3}
	number1, number2 = sample(array, match)
	fmt.Printf("numbers are: %v ,%v", number1, number2)
	fmt.Println()
}

