package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ** SOAL 1 *
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjangPersegiPanjangInt, _ := strconv.Atoi(panjangPersegiPanjang)
	lebarPersegiPanjangInt, _ := strconv.Atoi(lebarPersegiPanjang)
	alasSegitigaInt, _ := strconv.Atoi(alasSegitiga)
	tinggiSegitigaInt, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjangPersegiPanjangInt * lebarPersegiPanjangInt
	var kelilingPersegiPanjang int = 2 * (panjangPersegiPanjangInt + lebarPersegiPanjangInt)
	var luasSegitiga int = alasSegitigaInt * tinggiSegitigaInt / 2

	fmt.Printf("Luas Persegi Panjang: %d\n", luasPersegiPanjang)
	fmt.Printf("Keliling Persegi Panjang: %d\n", kelilingPersegiPanjang)
	fmt.Printf("Luas Segitiga: %d\n", luasSegitiga)


	// ** SOAL 2 *
	var nilaiJohn = 80
	var nilaiDoe = 50
	// nilaiJohn
	if nilaiJohn >= 80 {
		fmt.Println("Indeks nilai John: A")
	} else if nilaiJohn >= 70 {
		fmt.Println("Indeks nilai John: B")
	} else if nilaiJohn >= 60 {
		fmt.Println("Indeks nilai John: C")
	} else if nilaiJohn >= 50 {
		fmt.Println("Indeks nilai John: D")
	} else {
		fmt.Println("Indeks nilai John: E")
	}
	// nilaiDoe
	if nilaiDoe >= 80 {
		fmt.Println("Indeks nilai Doe: A")
	} else if nilaiDoe >= 70 {
		fmt.Println("Indeks nilai Doe: B")
	} else if nilaiDoe >= 60 {
		fmt.Println("Indeks nilai Doe: C")
	} else if nilaiDoe >= 50 {
		fmt.Println("Indeks nilai Doe: D")
	} else {
		fmt.Println("Indeks nilai Doe: E")
	}

	// ** SOAL 3 *
	var tanggal = 12
	var bulan = 1
	var tahun = 2003

	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan tidak valid"
	}

	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	// ** SOAL 4 *
	var tahunKelahiran = 2003

	var generasi string
	switch {
	case tahunKelahiran >= 1944 && tahunKelahiran <= 1964:
		generasi = "Baby boomer"
	case tahunKelahiran >= 1965 && tahunKelahiran <= 1979:
		generasi = "Generasi X"
	case tahunKelahiran >= 1980 && tahunKelahiran <= 1994:
		generasi = "Generasi Y (Millenials)"
	case tahunKelahiran >= 1995 && tahunKelahiran <= 2015:
		generasi = "Generasi Z"
	default:
		generasi = "Generasi belum terdefinisi"
	}

	fmt.Printf("Anda termasuk dalam generasi: %s\n", generasi)
}