package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// ** SOAL 1 *
	bootcamp := "Bootcamp"
	digital := "Digital"
	skill := "Skill"
	sanbercode := "Sanbercode"
	golang := "Golang"
	
	fmt.Println(bootcamp + " " + digital + " " + skill + " " + sanbercode + " " + golang) 

	// ** SOAL 2 *
	halo := "Halo Dunia"
	newHalo := strings.Replace(halo, "Dunia", "Golang", 1)

	fmt.Println(newHalo)

	// ** SOAL 3 *
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	
	resultSoal3 := kataPertama + " " + strings.Title(kataKedua) + " " + strings.Replace(kataKetiga, "a", "R", 1) + " " + strings.ToUpper(kataKeempat)

	fmt.Println(resultSoal3)

	// ** SOAL 4 *
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	num1, _ := strconv.Atoi(angkaPertama)
	num2, _ := strconv.Atoi(angkaKedua)
	num3, _ := strconv.Atoi(angkaKetiga)
	num4, _ := strconv.Atoi(angkaKeempat)

	total := num1 + num2 + num3 + num4

	fmt.Println(total)

	// ** SOAL 5 *
	kalimat := "halo halo bandung"
	angka := 2021

	sentenceFrag := strings.Split(kalimat, " ")
	sentenceFrag[0] = "Hi"
	sentenceFrag[1] = "Hi"
	sentenceFrag[2] = "bandung"

	newSentence := strings.Join(sentenceFrag, " ")
	newNum := strconv.Itoa(angka)

	resultSoal5 := `"` + newSentence + `"` + " - " + newNum

	fmt.Println(resultSoal5)

}