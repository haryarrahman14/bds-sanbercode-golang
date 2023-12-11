package main

import "fmt"

func note() {
	// Operator Aritmatika
	jumlah := 8 + 3
	fmt.Println("Penjumlahan:", jumlah) // hasilnya 11

	kurang := 8 - 3
	fmt.Println("Pengurangan:", kurang) // hasilnya 5

	kali := 8 * 3
	fmt.Println("Perkalian:", kali) // hasilnya 24

	bagi := 8 / 4
	fmt.Println("Pembagian:", bagi) // hasilnya 2

	modulus := 8 % 3
	fmt.Println("Modulus:", modulus) // hasilnya 2

	// Augmented Assignments
	var angka = 8
	fmt.Println("Nilai awal angka:", angka) // 8
	angka += 10
	fmt.Println("Setelah ditambah 10:", angka) // 18

	var angka2 = 5
	fmt.Println("Nilai awal angka2:", angka2) // 5
	angka2 += 5
	fmt.Println("Setelah ditambah 5:", angka2) // 10

	// Unary Operator
	angka3 := 8
	fmt.Println("Nilai awal angka3:", angka3) // 8
	angka3++
	fmt.Println("Setelah di-increment:", angka3) // 9

	angka4 := 5
	fmt.Println("Nilai awal angka4:", angka4) // 5
	angka4--
	fmt.Println("Setelah di-decrement:", angka4) // 4

	// Operator Perbandingan
	var angka5 = 8
	fmt.Println(angka5 > 5)  // true
	fmt.Println(angka5 < 5)  // false
	fmt.Println(angka5 >= 5) // true
	fmt.Println(angka5 <= 5) // false
	fmt.Println(angka5 == 5) // false
	fmt.Println(angka5 != 5) // true

	// Operator Logika
	var a, b, c, d = true, false, true, false
	fmt.Println(a && c)       // true
	fmt.Println(a && b)       // false
	fmt.Println(a || b)       // true
	fmt.Println(b || d)       // false
	fmt.Println(!b && !d)     // true
	fmt.Println(!a || b)      // false

	// Conditional (if, else if, else)
	var mood = "happy"
	if mood == "happy" {
		fmt.Println("Hari ini aku bahagia!")
	}

	var minimarketStatus = "open"
	if minimarketStatus == "open" {
		fmt.Println("Saya akan membeli telur dan buah")
	} else {
		fmt.Println("Minimarketnya tutup")
	}

	var minuteRemainingToOpen = 5
	if minimarketStatus == "open" {
		fmt.Println("Saya akan membeli telur dan buah")
		if minuteRemainingToOpen <= 5 {
			fmt.Println("Minimarket buka sebentar lagi, saya tungguin")
		} else {
			fmt.Println("Minimarket tutup, saya pulang lagi")
		}
	}

	// Switch-case
	var buttonPushed = 1
	switch buttonPushed {
	case 1:
		fmt.Println("Matikan TV!")
	case 2, 3, 4:
		fmt.Println("Turunkan volume TV!")
	default:
		fmt.Println("Tidak terjadi apa-apa")
	}

	// Switch-case dengan fallthrough
	var point = 6
	switch {
	case point == 8:
		fmt.Println("Perfect")
	case (point < 8) && (point > 3):
		fmt.Println("Awesome")
		fallthrough
	case point < 5:
		fmt.Println("You need to learn more")
	default:
		fmt.Println("Not bad")
		fmt.Println("You need to learn more")
	}
}
