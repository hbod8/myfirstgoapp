package main

import (
	"fmt"
)

func main() {

	// Basic stdout, Print, Println, Printf,
	fmt.Println("Hello, world!")

	// Basic stdin, Scanln and Pointers <3
	var userTalk string
	fmt.Print("C'mon don't be shy, what do you wannay say back to the world? >")
	fmt.Scanln(&userTalk)
	fmt.Printf("Wow, \"%s\", what brave first words!\n", userTalk)

	// Type inference using :=
	coolNumber := 5
	anotherCoolNumber := 7
	myCoolSum := coolNumber + anotherCoolNumber
	fmt.Printf("Mycool numbers: %d & %d, and their sum: %d\n", coolNumber, anotherCoolNumber, myCoolSum)

	// conditional execution
	fact := true
	if fact {
		fmt.Println("That's True!")
	} else {
		fmt.Println("That's False!")
	}

	//conditional loop
	number := 0
	for number < 3 {
		fmt.Println("the number is ", 3)
		number++
	}

	// for loop
	for myIterator := 0; myIterator < 5; myIterator++ {
		fmt.Println("the itererator is ", myIterator)
	}

	// arrays
	var myArray [5]int
	fmt.Println("my cute little array: ", myArray)
	mySecondArray := [5]int{8, 0, 0, 8, 5}
	fmt.Println("and his friend: ", mySecondArray)

	// slices (slow, array copy operations)
	var mySlice []int
	fmt.Println("look at this empty slice: ", mySlice)
	mySlice = append(mySlice, 1)
	fmt.Println("awww, now it's got a 1, it's 1 and only: ", mySlice)
}