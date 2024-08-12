package main

import "fmt"

func () { 
	
	// string-> []byte
	// in ASCII -> lowercase and uppercase letter are diffrent
	// e.g. a and A 

	word := "academy"

	for _, char := range word {
		fmt.Println(char)
		

	}
}