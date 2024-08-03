package main

import "fmt"

func main() {

	// slice works like a array, but more flexible size

	// approach 1
	var books []string

	// add item -> append()
	books = append(books, " Data Science")
	books = append(books, "Algorithm", "Digital marketing", "full stack development")

	// for range
	for _, books := range books {
		fmt.Println(books)
	}
}
