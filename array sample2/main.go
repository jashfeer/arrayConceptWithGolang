package main

import "fmt"

func sample(arr [10]int, target int) (int, int) {
	mymap := map[int]struct{}{}
	for i := 0; i < len(arr); i++ {
		match := target - arr[i]
		_, ok := mymap[match]
		if ok == true {
			return arr[i], match
		} else {
			mymap[arr[i]] = struct{}{}
		}
	}
	return 0, 0
}


//o(n)ST

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 2, 3, 1, 4}
	value := 10
	number1, number2 := sample(array, value)
	fmt.Printf("Two values are %v and %v", number1, number2)
}
