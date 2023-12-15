package main

import "fmt"

// Contoh 1: Function Sederhana
func printHello() {
	fmt.Println("Hello")
}

// Contoh 2: Function dengan Parameter
func printAngka(angka1 int, angka2 int) {
	fmt.Println("angka pertama", angka1)
	fmt.Println("angka kedua", angka2)
}

// Contoh 3: Function dengan Return Value
func introduction(name string) string {
	return "Hello My Name Is " + name
}

// Contoh 4: Function dengan Multiple Return Value
func introducenote(firstName string, lastName string) (string, string) {
	introFirstName := "Hello My First Name Is " + firstName
	introLastName := "Hello My Last Name Is " + lastName
	return introFirstName, introLastName
}

// Contoh 5: Function dengan Predefined Return Value
func tambahAngka(firstNumber int, lastNumber int) (jumlah int) {
	jumlah = firstNumber + lastNumber
	return
}

func tampilkanKataDanAngka() (firstWord, secondWord string, number int) {
	firstWord = "Halo"
	secondWord = "Dunia"
	number = 10
	return
}

// Contoh 6: Variadic Function
func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Contoh 7: Function sebagai Parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// Contoh 8: Closure Function
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

func note() {
	// Contoh 1: Memanggil Function Sederhana
	printHello()

	// Contoh 2: Memanggil Function dengan Parameter
	printAngka(1, 2)

	// Contoh 3: Memanggil Function dengan Return Value
	fmt.Println(introduction("John"))

	// Contoh 4: Memanggil Function dengan Multiple Return Value
	firstName, lastName := introducenote("John", "Doe")
	fmt.Println(firstName, lastName)

	// Contoh 5: Memanggil Function dengan Predefined Return Value
	hasil := tambahAngka(4, 5)
	fmt.Println(hasil)

	kataPertama, kataKedua, angka := tampilkanKataDanAngka()
	fmt.Println(kataPertama, kataKedua, angka)

	// Contoh 6: Memanggil Variadic Function
	var total = sum(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	fmt.Println(total)

	// Contoh 7: Memanggil Function sebagai Parameter
	sayHelloWithFilter("John", func(name string) string {
		return "Mr. " + name
	})

	sayHelloWithFilter("Kasar", func(name string) string {
		if name == "Kasar" {
			return "..."
		} else {
			return name
		}
	})

	// Contoh 8: Memanggil Closure Function
	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()
	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}
