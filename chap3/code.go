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

func main() {
	fmt.Println(concat("1", "2"))
	// in go we pass variables by value, not by reference except for a few data types not yet seen
	x := 5
	x = increment((x))
	fmt.Println(x)

	s1, _ := getNames() // ignore return value, the compiler will remove this from our code, useful since we cannot have any unused variables
	fmt.Println(s1)
	// named return values 1:06:27
}
