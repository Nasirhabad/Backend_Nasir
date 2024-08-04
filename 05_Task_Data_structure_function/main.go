package main

import "fmt"

func main() {
	// array-> store many item in espacific size (fixed size)
	//approach one
	// var numbers [5]int

	//inset items
	// numbers[0] = 12
	// numbers[1] = 14
	// numbers[2] = 46
	// numbers[3] = 50

	// // approach 1 for accessing array -> regular for loop
	// // more costomizable (epecify which index)
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])

	// }
	// // 2nd approad for accessing  array -> for range
	// for idx, val := range numbers {
	// 	fmt.Println("the index: ", idx, "the value: ", val)
	// }

	// slice
	var books []string

	// add item --> append ()

	books = append(books, "backend")
	books = append(books, "frontend", "full stack", "web dev")

	// // for range
	// for idx, books := range books {
	// 	fmt.Println("index: ", idx, "value", books)
	// }
	// for loop -> reversed

	for i := len(books) - 1; i >= 0; i-- {
		fmt.Println(books[i])
	}

}
