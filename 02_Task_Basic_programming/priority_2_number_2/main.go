package main

import (
	"fmt"
)

// // func main() {
// // 	number := 70
// // 	fmt.Println(generatePattern(number))

// // 	number = 35
// // 	fmt.Println(generatePattern(number))
// // }

// // func generatePattern(number int) string {
// // 	factors := findFactors(number)
// // 	pattern := ""
// // 	for _, factor := range factors {
// // 		if factor%2 == 0 {
// // 			pattern += "I"
// // 		} else {
// // 			pattern += "O"
// // 		}
// // 	}
// // 	return pattern
// // }

// // func findFactors(number int) []int {
// // 	factors := []int{}
// // 	for i := 1; i <= number; i++ {
// // 		if number%i == 0 {
// // 			factors = append(factors, i)
// // 		}
// // 	}
// // 	return factors
// }

func main() {
	var number int

	fmt.Print("enter a number: ")
	fmt.Scan(&number)

	// find the number factors
	for i := 1; i <= number; i++ {
		if number%i == 0 {

			fmt.Println(i)
			if i%2 == 0 {
				fmt.Println("I")
			} else {
				fmt.Println("O")
			}

		}
	}
}
