package main

import "fmt"

type person struct {
    name string
    age  int
}

type student struct {
    name  string
    grade int
}

type studentLiterals struct {
	name  string
    grade int
}

type studentEmbedded struct {
	name  string
    grade int
	person
}

type studentNested struct {
	name  string
    grade int
	person
}

func note() {
    // Deklarasi Struct
    var john student
    john.name = "john doe"
    john.grade = 2

    fmt.Println("name  :", john.name)
    fmt.Println("grade :", john.grade)

    // Struct Literals
    // cara pertama
    var jack = studentLiterals{}
    jack.name = "jack"
    jack.grade = 3

    // cara kedua tetapi isinya harus berurutan
    var jane = studentLiterals{"jane", 4}

    // cara ketiga dengan nama property tetapi tidak harus berurutan
    var jill = studentLiterals{name: "jill", grade: 5}

    fmt.Println("student 1 :", jack.name)
    fmt.Println("student 2 :", jane.name)
    fmt.Println("student 3 :", jill.name)

    // Embedded Struct
    // contoh 1
    var embeddedJohn = studentEmbedded{}
    embeddedJohn.name = "john"
    embeddedJohn.age = 21
    embeddedJohn.grade = 2

    fmt.Println("name  :", embeddedJohn.name)
    fmt.Println("age   :", embeddedJohn.age)
    fmt.Println("grade :", embeddedJohn.grade)

    // contoh 2
    var doeData = person{name: "doe", age: 21}
    var embeddedDoe = studentEmbedded{person: doeData, grade: 2}

    fmt.Println("name  :", embeddedDoe.person.name)
    fmt.Println("age   :", embeddedDoe.person.age)
    fmt.Println("grade :", embeddedDoe.grade)

    // Anonymous Struct
    // anonymous struct tanpa pengisian property
    var anonymousJohn = struct {
        student
        address string
    }{}

    // anonymous struct dengan pengisian property
    var anonymousDoe = struct {
        student
        address string
    }{
        student: student{name: "doe", grade: 2},
        address: "123 Main St",
    }

    fmt.Println("name  :", anonymousJohn.student.name)
    fmt.Println("grade :", anonymousJohn.student.grade)

    fmt.Println("name  :", anonymousDoe.student.name)
    fmt.Println("grade :", anonymousDoe.student.grade)
    fmt.Println("address :", anonymousDoe.address)

    // Nested Struct
    var nestedJohn = studentNested{}
    nestedJohn.person.name = "john"
    nestedJohn.person.age = 21
    nestedJohn.grade = 2

    fmt.Println("name  :", nestedJohn.person.name)
    fmt.Println("age   :", nestedJohn.person.age)
    fmt.Println("grade :", nestedJohn.grade)

    // Method
    var johnWick = student{"john wick", 21}
    johnWick.sayHello()
}

// Method
func (s student) sayHello() {
    fmt.Println("halo", s.name)
}
