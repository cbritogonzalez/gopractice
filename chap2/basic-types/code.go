package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	test := "Hi"
	fmt.Printf(test+" is of type %T "+"\n", test)
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
	// same line declaration
	rate, message := .23, "the avg is "
	fmt.Printf("The rate is %v and %q\n", rate, message)
	// int, uint, float64, string, bool are the most comun, if needed we may look for another type anotation to optimize
	// casting
	var fromOneType float64 = 2.0
	fmt.Printf("type of this is %T, %f \n", fromOneType, fromOneType)
	fmt.Printf("to this type  %T , %v\n", int(fromOneType), int(fromOneType))
	// we also have const, does not support short declaration, known at compile time, we can create other const value based from other const values
	// we can have
	const constantValue = "Carlos"
	const constantValue2 = "Brito"
	const name = constantValue + " " + constantValue2 + "\n"
	fmt.Printf(name)
	// string formatting fmt.Printf
	// %v is the go syntax representation of a value, use if unsure what to use, but it is better to use a specific indicator
	// %s interpolate a string
	// %d interpolate an integer in decimal form
	// %f interpolate a decimal
	// check fmt package for more details
	msg := fmt.Sprintf("Hi %s, your rate is %.1f", name, rate) // use this to format and return a string
	println(msg)
	// conditionals, check for boolean conditionals on page
	// == , != , < , > , <= , >=
	// rate = 2
	if rate < 1 {
		fmt.Println("if")
	} else if rate > 1 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
	if length := len(name); length > 1 { //length is only accessible within the scope of the if block
		fmt.Println("length test")
	}

}
