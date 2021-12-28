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
	fmt.Printf("My cool numbers: %d & %d, and their sum: %d\n", coolNumber, anotherCoolNumber, myCoolSum)

	// conditional execution
	fact := true
	if fact {
		fmt.Println("That's True!")
	} else {
		fmt.Println("That's False!")
	}

	// conditional loop
	number := 0
	for number < 3 {
		fmt.Println("the number is", number)
		number++
	}

	// for loop
	for myIterator := 0; myIterator < 5; myIterator++ {
		fmt.Println("the itererator is", myIterator)
	}

	// arrays
	var myArray [5]int
	fmt.Println("my cute little array:", myArray)
	mySecondArray := [5]int{8, 0, 0, 8, 5}
	fmt.Println("and his friend:", mySecondArray)

	// multidimensional arrays
	var myTwoDArray [3][3]int
	fmt.Println("yo, this dude braking the fourth wall:", myTwoDArray)
	var myFiveDArray [2][2][2][2][2]int
	fmt.Println("we we're too busy asking if we could, we didn't ask if we should:", myFiveDArray)

	// slices (slow, array copy operations)
	var mySlice []int // also make zero-valued ones using make([]int, 5), additionally you can copy slices using copy(a, b)
	fmt.Println("look at this empty slice: ", mySlice)
	mySlice = append(mySlice, 1) // also handles multiple appends as additional arguments
	fmt.Println("awww, now it's got a 1, it's 1 and only: ", mySlice) // we can extract ranges using mySlice[low:high] (not incl. high)
	// this means delete can be made using append(mySlice[:i], mySlice[i+1:]...)

	// map & deletions
	myMap := make(map[string]int)
	myMap["My Favorite Number"] = 2
	myMap["My Second Favorite Number"] = 4
	myMap["My SSN"] = 1234567890
	delete("My SSN")
	fmt.Println(myMap)

	// set loop


	// functions
	// NOTE: variadic functions can be made using func() (...int) and if you are passing a slice
	// and you can send the elements as separate arguments as: myFunction(mySlice...)
}