package helper

import (
	"errors"
	"fmt"
	"strconv"
)

func KelilingSegitigaSamaSisi(sisi int, tampilkanKalimat bool) (string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	keliling := 3 * sisi

	if sisi != 0 {
		if tampilkanKalimat {
			// jika parameter kedua bernilai true maka tampilkan kalimat (asumsi sisinya 2) "keliling segitiga sama sisinya dengan sisi 2 cm adalah 6 cm"
			return "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(sisi) + "cm adalah " + strconv.Itoa(keliling) + " cm"
		} else {
			// jika parameter kedua bernilai false maka tampilkan hanya hasil string angka saja (misal "6")	
			return strconv.Itoa(keliling)
		}
	}

	// jika parameter pertama 0 dan parameter kedua bernilai true gunakan error dengan message "Maaf anda belum menginput sisi dari segitiga sama sisi"
	if tampilkanKalimat {
		errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		return ""
	} else {
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		panic(fmt.Errorf("%v\n%s", err, "panic: Maaf anda belum menginput sisi dari segitiga sama sisi"))
	}
}