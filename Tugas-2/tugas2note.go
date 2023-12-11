package main

import (
	"fmt"
	"strconv"
	"strings"
)

func note()  {
	// Introduction to Golang
	fmt.Println("Hello, Gophers! Let's explore Golang.")

	// Installation of Go
	// Note: Adjust the version number accordingly
	fmt.Println("Please follow the installation steps mentioned in the article.")

	// Creating the First Go Program
	fmt.Println("Creating the first Go program:")
	fmt.Println("package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"Hello World\")\n}")

	// Go Modules
	fmt.Println("Initializing a new Go module:")
	fmt.Println("go mod init <nama-project>")
	fmt.Println("Example: go mod init firstProject")

	// Variable and Constant
	var message string = "Hello, Go!"
	fmt.Println("Example of variable declaration and initialization:")
	fmt.Println("var message string = \"Hello, Go!\"")
	fmt.Println("Message:", message)

	const pi = 3.14
	fmt.Println("\nExample of constant declaration:")
	fmt.Println("const pi = 3.14")
	fmt.Println("Value of pi:", pi)

	// Data Types
	// Numeric Data Types
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	fmt.Printf("\nExample of numeric data types:\nPositive Number: %d\nNegative Number: %d\n", positiveNumber, negativeNumber)

	// Boolean Data Type
	var exist bool = true
	fmt.Printf("\nExample of boolean data type:\nExist? %t\n", exist)

	// String Data Type
	var str = "Go Programming"
	fmt.Printf("\nExample of string data type:\nString: %s\n", str)

	// Nil and Zero Value
	var emptyString string
	var isTrue bool
	var zero int
	fmt.Printf("\nExample of nil and zero value:\nEmpty String: \"%s\"\nBoolean: %t\nZero: %d\n", emptyString, isTrue, zero)

	// Type Conversion (Casting)
	var num int = 42
	var numStr string = strconv.Itoa(num)
	fmt.Printf("\nExample of type conversion (int to string):\nConverted String: %s\n", numStr)

	// String Functions from strings package
	var sentence = "Go is awesome!"
	fmt.Printf("\nExample of string functions from strings package:\nOriginal Sentence: %s\n", sentence)
	fmt.Printf("Index of 'is': %d\n", strings.Index(sentence, "is"))
	fmt.Printf("After Replace 'awesome' with 'fantastic': %s\n", strings.Replace(sentence, "awesome", "fantastic", 1))
	fmt.Printf("Repeated 'Go': %s\n", strings.Repeat("Go ", 3))

	// String Conversion Functions from strconv package
	var numericStr = "42"
	numeric, _ := strconv.Atoi(numericStr)
	fmt.Printf("\nExample of string conversion functions from strconv package:\nConverted Numeric Value: %d\n", numeric)

	// Additional functions from strconv package
	var floatStr = "3.14"
	floatNum, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("Parsed Float Value: %f\n", floatNum)

	var boolStr = "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("Parsed Boolean Value: %t\n", boolValue)
}
