package helper

import "fmt"

func TambahAngka(nilai int, angka *int) {
	*angka += nilai
}

func CetakAngka(angka *int) {
	fmt.Println(*angka)
}