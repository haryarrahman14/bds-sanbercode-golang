package main

import (
	"fmt"
	"math"
	"strings"
)

// ** SOAL 1 *
type segitigaSamaSisi struct{
	alas, tinggi float64
}

type persegiPanjang struct{
	panjang, lebar float64
}

type tabung struct{
	jariJari, tinggi float64
}

type balok struct{
	panjang, lebar, tinggi float64
}

type hitungBangunDatar interface{
	luas() 		float64
	keliling()	float64
}

type hitungBangunRuang interface{
	volume()		float64
	luasPermukaan() float64
}

// ** IMPLEMENTASI SOAL 1 *

// ** BANGUN DATAR *
func (s segitigaSamaSisi) luas() float64 {
	return float64(s.alas * s.alas) / 2
}

func (s segitigaSamaSisi) keliling() float64 {
	return 3 * float64(s.alas)
}

func (pp persegiPanjang) luas() float64 {
	return float64(pp.panjang * pp.lebar)
}

func (pp persegiPanjang) keliling() float64 {
	return 2 * float64(pp.panjang + pp.lebar)
}

// ** BANGUN RUANG *
func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * float64(t.tinggi)
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (float64(t.jariJari) + t.tinggi)
}

// Implementasi interface hitungBangunRuang untuk balok
func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64((b.panjang*b.lebar) + (b.panjang*b.tinggi) + (b.lebar*b.tinggi))
}

// ** SOAL 2 *
type phone struct{
	name, brand	string
	year		int
	colors		[]string
 }

type displayPhoneInfo interface {
	getInfo() string
}

func (p phone) getInfo() string {
	colorStr := strings.Join(p.colors, ", ")
	info := fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.name, p.brand, p.year, colorStr)
	return info
}

func getDisplayPhoneInfo(d displayPhoneInfo) {
	info := d.getInfo()

	fmt.Println(info)
}

// ** SOAL 3 *
func luasPersegi(sisi int, showDetail bool) interface{} {
	if sisi == 0 {
		if showDetail {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi
	if showDetail {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}

// ** SOAL 4 *
func sum(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

// Fungsi untuk mengubah slice angka menjadi format "a + b + c + ..."
func formatNumbers(numbers ...[]int) string {
	var formattedNumbers []string
	for _, nums := range numbers {
		formattedNumbers = append(formattedNumbers, formatSlice(nums)...)
	}
	return strings.Join(formattedNumbers, " + ")
}

// Fungsi untuk mengubah slice angka menjadi slice string
func formatSlice(numbers []int) []string {
	var formattedNumbers []string
	for _, num := range numbers {
		formattedNumbers = append(formattedNumbers, fmt.Sprint(num))
	}
	return formattedNumbers
}


func main() {
	// ** SOAL 1 *
	// Menggunakan interface hitungBangunDatar untuk segitigaSamaSisi
	var bangunDatar1 hitungBangunDatar = segitigaSamaSisi{alas: 5, tinggi: 7}
	fmt.Println("===== Segitiga Sama Sisi")
	fmt.Println("Luas:", bangunDatar1.luas())
	fmt.Println("Keliling:", bangunDatar1.keliling())
	fmt.Println()

	// Menggunakan interface hitungBangunDatar untuk persegiPanjang
	var bangunDatar2 hitungBangunDatar = persegiPanjang{panjang: 4, lebar: 6}
	fmt.Println("===== Persegi Panjang")
	fmt.Println("Luas:", bangunDatar2.luas())
	fmt.Println("Keliling:", bangunDatar2.keliling())
	fmt.Println()

	// Menggunakan interface hitungBangunRuang untuk tabung
	var bangunRuang1 hitungBangunRuang = tabung{jariJari: 3.0, tinggi: 8}
	fmt.Println("===== Tabung (Silinder)")
	fmt.Println("Volume:", bangunRuang1.volume())
	fmt.Println("Luas Permukaan:", bangunRuang1.luasPermukaan())
	fmt.Println()

	// Menggunakan interface hitungBangunRuang untuk balok
	var bangunRuang2 hitungBangunRuang = balok{panjang: 2, lebar: 3, tinggi: 4}
	fmt.Println("===== Balok")
	fmt.Println("Volume:", bangunRuang2.volume())
	fmt.Println("Luas Permukaan:", bangunRuang2.luasPermukaan())
	fmt.Println()

	// ** SOAL 2 *
	myPhone := phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	getDisplayPhoneInfo(myPhone)
	fmt.Println()

	// ** SOAL 3 *
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// ** SOAL 4 *
	var prefix interface{}= "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6,8}
	var kumpulanAngkaKedua interface{} = []int{12,14}

	// tulis jawaban anda disini
	angkaPertama, ok1 := kumpulanAngkaPertama.([]int)
	angkaKedua, ok2 := kumpulanAngkaKedua.([]int)

	if ok1 && ok2 {
		total := sum(angkaPertama) + sum(angkaKedua)
		output := fmt.Sprintf("%s%s = %d", prefix, formatNumbers(angkaPertama, angkaKedua), total)
		fmt.Println(output)
	} else {
		fmt.Println("Invalid data type")
	}
	

}