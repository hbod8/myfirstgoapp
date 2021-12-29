package main

import (
	"fmt"
	"sync"
	"time"
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
	delete(myMap, "My SSN")
	fmt.Println(myMap)

	// loop over a set
	for myIndex, myValue := range myArray { // also works for value only (for v := range myArray) and maps
		fmt.Println("look, i found a", myValue, "at position", myIndex)
	}

	// functions
	fmt.Println("let's ask my friend, sum, to add these numbers, he says:", sum(9, 10))

	// multi return functions
	mul, div := multiplyAndDivide(8, 4) // if i did not need a return, i can simply sure _ instead of a variable name
	fmt.Println("my other friend, multiplyAndDivide is very good at giving me all the info:", mul, "and", div)

	// varadic functions
	fmt.Println("I can send this function soooo many arguments:", sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // and you can send the elements as separate arguments as: myFunction(mySlice...)

	// closures & returned functinos
	anotherFunction := myNumberFunction()
	fmt.Println(anotherFunction())
	fmt.Println(anotherFunction())
	fmt.Println(anotherFunction())

	// inline functions
	printSilly := func() {
		fmt.Println("silly")
	}
	printSilly()

	// structs
	bigBoy := train{1200, 24, 1000000, "UP 4014"}
	fmt.Println("Look at this train:", bigBoy)

	// methods
	fmt.Println("it can pull like:", bigBoy.traction())

	// interfaces are supported and look like:
	// type vechicle interface{
	// 	getHorepower() int
	// }
	// but I dont use interfaces that much

	// concurrency & parraelleism & wait groups
	fmt.Println("Let's race and see who can count to 20 the fastest!")
	var myWaitGroup sync.WaitGroup
	racer := func() {
		for i := 1; i <= 20; i++ {
			fmt.Print(i, " ")
			time.Sleep(100 * time.Millisecond)
		}
	}
	startRacer := func() {
		racer()
		myWaitGroup.Done()
	}
	myWaitGroup.Add(3)
	go startRacer()
	go startRacer()
	go startRacer()
	myWaitGroup.Wait()
	fmt.Println()

	// channels
	myChannel := make(chan int)
	producer := func(count int, c chan int) {
		for i := 0; i < count; i++ {
			fmt.Println("I'm making a", i)
			fmt.Println("I'm sending a", i)
			c <- i
		}
	}
	consumer := func(count int, c chan int) {
		for i := 0; i < count; i++ {
			food := <- c
			fmt.Println("I got a", food)
		}
	}
	myWaitGroup.Add(2)
	go func() {
		producer(5, myChannel)
		myWaitGroup.Done()
	}()
	go func() {
		consumer(5, myChannel)
		myWaitGroup.Done()
	}()
	myWaitGroup.Wait()
}

// functions
func sum(a int, b int) int {
	return a + b
}

// multiple return functions
func multiplyAndDivide(a int, b int) (int, int) {
	return a * b, a / b
}

// varadic functions
func sumAll(n ...int) int {
	sum := 0
	for v := range n {
		sum += v
	}
	return sum
}

// closures
func myNumberFunction() func() int {
	mySuperSpecialNumber := 10
	return func() int {
		mySuperSpecialNumber++
		return mySuperSpecialNumber
	}
}

// structs
type train struct { // you can also embeed another struct within a struct (think JSON)
	power int
	wheels int
	weight int
	name string
}

// methods
func (t train) traction() int {
	return t.weight * t.power
}