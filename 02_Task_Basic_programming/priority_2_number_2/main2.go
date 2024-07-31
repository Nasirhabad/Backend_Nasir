package main

import (
	"fmt"
)

func main() {
	var budget, duration, difficulty int

	fmt.Print("Enter a budget: ")
	fmt.Scan(&budget)

	fmt.Print("Enter a duration: ")
	fmt.Scan(&duration)

	fmt.Print("Enter a difficulty: ")
	fmt.Scan(&difficulty)

	var budgetScore, durationScore, diffScore int

	// Assign the budget score
	if budget >= 0 && budget <= 50 {
		budgetScore = 4
	} else if budget >= 51 && budget <= 80 {
		budgetScore = 3
	} else if budget >= 81 && budget <= 100 {
		budgetScore = 2
	} else if budget > 100 {
		budgetScore = 1
	}

	// Assign the duration score
	if duration >= 0 && duration <= 20 {
		durationScore = 10
	} else if duration >= 21 && duration <= 30 {
		durationScore = 5
	} else if duration >= 31 && duration <= 50 {
		durationScore = 2
	} else if duration > 50 {
		durationScore = 1
	}

	// Assign the difficulty score
	if difficulty >= 0 && difficulty <= 3 {
		diffScore = 10
	} else if difficulty >= 4 && difficulty <= 6 {
		diffScore = 5
	} else if difficulty >= 7 && difficulty <= 10 { // Fix the range
		diffScore = 1
	} else if difficulty > 10 {
		diffScore = 0
	}

	// Calculate the priority score
	var priorityScore int = budgetScore + durationScore + diffScore

	// Check the priority of the project
	if priorityScore >= 17 && priorityScore <= 24 {
		fmt.Println("High")
	} else if priorityScore >= 10 && priorityScore <= 16 {
		fmt.Println("Medium")
	} else if priorityScore >= 3 && priorityScore <= 9 {
		fmt.Println("Low")
	} else if priorityScore < 3 { // Fix the condition
		fmt.Println("Impossible")
	}
}
