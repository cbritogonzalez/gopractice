package main

import "fmt"

func concat(s1 string, s2 string) string { // function receives two strings and returns a string
	return s1 + s2
}
func concat2(s1, s2 string) string { // function can also be declared like this and the type of the arguments are the same in this case
	return s1 + s2
}

// func increment(x int) {
// 	x++
// }

func increment(x int) int {
	x++
	return x
}

func getNames() (string, string) { // multiple returning values, wrap
	return "John", "Doe"
}

// func getCoords() (x, y int) { // naked return statements, returns x, y
// 	return
// }

// func getCoords() (int, int){
// 	var x int
// 	var y int
// 	return x,y
// }

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) { // implicit return, could be explicit if we put the three variables in order in the return like python
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

func main() {
	fmt.Println(concat("1", "2"))
	// in go we pass variables by value, not by reference except for a few data types not yet seen
	x := 5
	x = increment((x))
	fmt.Println(x)

	s1, _ := getNames() // ignore return value, the compiler will remove this from our code, useful since we cannot have any unused variables
	fmt.Println(s1)
	// named return values
	age := 18
	fmt.Println(yearsUntilEvents(age))

	// named return values 1:15:00
}
