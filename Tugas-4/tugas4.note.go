package main

import "fmt"

func note() {
	// Array
	var names [4]string
	names[0] = "John"
	names[1] = "Doe"
	names[2] = "Frank"
	names[3] = "Jack"

	fmt.Println(names[0], names[1], names[2], names[3])

	// Inisialisasi Nilai Awal Array
	var fruits = [4]string{"John", "Doe", "Frank", "Jack"}
	fmt.Println(fruits[0], fruits[1], fruits[2], fruits[3])

	// Array Multidimensi
	var numbers1 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// Slice
	var fruitsSlice = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruitsSlice[0])

	// Slice Merupakan Tipe Data Reference
	var aFruits = fruitsSlice[0:2]
	var bFruits = fruitsSlice[1:3]
	fmt.Println("aFruits", aFruits)
	fmt.Println("bFruits", bFruits)

	// Fungsi len() dan cap() pada Slice
	fmt.Println("len:", len(fruitsSlice))
	fmt.Println("cap:", cap(fruitsSlice))
	fmt.Println("len(aFruits):", len(aFruits))
	fmt.Println("cap(aFruits):", cap(aFruits))
	fmt.Println("len(bFruits):", len(bFruits))
	fmt.Println("cap(bFruits):", cap(bFruits))

	// Fungsi append() pada Slice
	var cFruits = append(bFruits, "pineapple")
	fmt.Println("cFruits", cFruits)
	fmt.Println("fruitsSlice", fruitsSlice)

	// Fungsi copy() pada Slice
	var dFruits = []string{"potato", "potato", "potato"}
	var eFruits = fruitsSlice[0:2]
	copy(dFruits, eFruits)
	fmt.Println("dFruits", dFruits)

	// Pengaksesan Elemen Slice Dengan 3 Indeks
	var fFruits = fruitsSlice[0:2:2]
	fmt.Println("fFruits", fFruits)

	// Map
	var chicken = map[string]int{"januari": 50, "februari": 40}
	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	// Deteksi Keberadaan Item Map
	var value, isExist = chicken["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// Looping dengan for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// Looping dengan for dan argumen hanya kondisi
	var j = 0
	for j < 5 {
		fmt.Println("Angka", j)
		j++
	}

	// Looping tanpa argumen
	var k = 0
	for {
		fmt.Println("Angka", k)
		k++
		if k == 5 {
			break
		}
	}

	// Perulangan Elemen Array Menggunakan for - range
	var fruitsArray = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruitsArray {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// Penggunaan Variabel Underscore _ Dalam for - range
	var fruitsRange = [4]string{"apple", "grape", "banana", "melon"}
	for _, fruit := range fruitsRange {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// Penggunaan break & continue
	for l := 1; l <= 10; l++ {
		if l%2 == 1 {
			continue
		}

		if l > 8 {
			break
		}

		fmt.Println("Angka", l)
	}
}
