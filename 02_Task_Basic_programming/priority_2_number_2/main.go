// // package main

// // import (
// // 	"fmt"
// // )

// // // // func main() {
// // // // 	number := 70
// // // // 	fmt.Println(generatePattern(number))

// // // // 	number = 35
// // // // 	fmt.Println(generatePattern(number))
// // // // }

// // // // func generatePattern(number int) string {
// // // // 	factors := findFactors(number)
// // // // 	pattern := ""
// // // // 	for _, factor := range factors {
// // // // 		if factor%2 == 0 {
// // // // 			pattern += "I"
// // // // 		} else {
// // // // 			pattern += "O"
// // // // 		}
// // // // 	}
// // // // 	return pattern
// // // // }

// // // // func findFactors(number int) []int {
// // // // 	factors := []int{}
// // // // 	for i := 1; i <= number; i++ {
// // // // 		if number%i == 0 {
// // // // 			factors = append(factors, i)
// // // // 		}
// // // // 	}
// // // // 	return factors
// // // }

// // func main() {
// // 	var number int

// // 	fmt.Print("enter a number: ")
// // 	fmt.Scan(&number)

// // 	// find the number factors
// // 	for i := 1; i <= number; i++ {
// // 		if number%i == 0 {

// // 			if i%2 == 0 {

// // 				fmt.Print("I")
// // 			} else {
// // 				fmt.Print("O")
// // 			}

// // 		}
// // 	}
// // }

// package main

// import (
// 	"fmt"
// )
// func main() {
// 	var budget, duration, difficulty int

// 	fmt.Print("enter a budget: ")
// 	fmt.Scan(&budget)

// 	fmt.Print("enter a duration: ")
// 	fmt.Scan(&duration)

// 	fmt.Print("enter a dificulty: ")
// 	fmt.Scan(&difficulty)

// 	var budgetScore, durationScore, diffScore int

// 	// TODO: assign the budget score
// 	if budget >= 0 && budget <= 50 {
// 		budgetSore = 4
// 	} else if budget >= 51 && budget <= 80 {
// 		bubudgetScore = 3

// 	} else if budget >= 81 &&  budget <= 100 {
// 		budgetScore = 2
// 	} else if budget > 100{
// 		budgetScore = 1
// 	}
// 	//TODO: assign the duration score

// 	if duration >= 0 && duration <= 20 {
// 		durationScore = 10
// 	} else if duration >= 21 && duration <= 30  {
// 		durartionSore = 5
// 	} else if duration >= 31 &&  duration <= 50 {
// 		durationScore= 2
// 	} else if duration > 50 {
// 		durationScore = 1
// 	}

// 	//TODO: assign the difficulty score

// 	if difficulty >= 0 && difficulty <= 3 {
// 		diffScore = 10
// 	} else if difficulty >= 4 && difficulty <= 6 {
// 		diffScore = 5
// 	} else if difficulty >= 8 &&  difficulty <= 10 {
// 		diffScore = 1
// 	} else if difficulty >10  {
// 		diffScore = 0
// 	}

// 	var priorityScore int budgetScore + durationScore + diffScore

// 	// check  priority of the project

// 	if priorityScore >= 17 && priorityScore <= 24 {
// 	fmt.Println("high")
// 	} else if priorityScore >= 10 && priorityScore <= 16 {
// 		fmt.Println("medium")
// 	} else if priorityScore >= 3 && priorityScore <= 9 {
// 		fmt.Println("low")
// 	} else if priorityScore < 2  {
// 		fmt.Println("impossiple")
// 	}

// }
