package main

import (
	"bds-sanbercode-golang/Tugas-9/controller"
	"bds-sanbercode-golang/Tugas-9/helper"
	"bds-sanbercode-golang/Tugas-9/model"
	"fmt"
)

func main() {
	// ** SOAL 1 *
	// Menggunakan interface hitungBangunDatar untuk segitigaSamaSisi
	var bangunDatar1 model.HitungBangunDatar = model.SegitigaSamaSisi{Alas: 5, Tinggi: 7}
	fmt.Println("===== Segitiga Sama Sisi")
	fmt.Println("Luas:", bangunDatar1.Luas())
	fmt.Println("Keliling:", bangunDatar1.Keliling())
	fmt.Println()

	// Menggunakan interface hitungBangunDatar untuk persegiPanjang
	var bangunDatar2 model.HitungBangunDatar = model.PersegiPanjang{Panjang: 4, Lebar: 6}
	fmt.Println("===== Persegi Panjang")
	fmt.Println("Luas:", bangunDatar2.Luas())
	fmt.Println("Keliling:", bangunDatar2.Keliling())
	fmt.Println()

	// Menggunakan interface hitungBangunRuang untuk tabung
	var bangunRuang1 model.HitungBangunRuang = model.Tabung{JariJari: 3.0, Tinggi: 8}
	fmt.Println("===== Tabung (Silinder)")
	fmt.Println("Volume:", bangunRuang1.Volume())
	fmt.Println("Luas Permukaan:", bangunRuang1.LuasPermukaan())
	fmt.Println()

	// Menggunakan interface hitungBangunRuang untuk balok
	var bangunRuang2 model.HitungBangunRuang = model.Balok{Panjang: 2, Lebar: 3, Tinggi: 4}
	fmt.Println("===== Balok")
	fmt.Println("Volume:", bangunRuang2.Volume())
	fmt.Println("Luas Permukaan:", bangunRuang2.LuasPermukaan())
	fmt.Println()

	// ** SOAL 2 *
	myPhone := model.Phone{
		Name:	"Samsung Galaxy Note 20",
		Brand:	"Samsung",
		Year:	2020,
		Colors:	[]string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	controller.GetDisplayPhoneInfo(myPhone)
	fmt.Println()

	// ** SOAL 3 *
	fmt.Println(helper.LuasPersegi(4, true))
	fmt.Println(helper.LuasPersegi(8, false))
	fmt.Println(helper.LuasPersegi(0, true))
	fmt.Println(helper.LuasPersegi(0, false))

	// ** SOAL 4 *
	var prefix interface{}= "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6,8}
	var kumpulanAngkaKedua interface{} = []int{12,14}

	// tulis jawaban anda disini
	angkaPertama, ok1 := kumpulanAngkaPertama.([]int)
	angkaKedua, ok2 := kumpulanAngkaKedua.([]int)

	if ok1 && ok2 {
		total := helper.Sum(angkaPertama) + helper.Sum(angkaKedua)
		output := fmt.Sprintf("%s%s = %d", prefix, helper.FormatNumbers(angkaPertama, angkaKedua), total)
		fmt.Println(output)
	} else {
		fmt.Println("Invalid data type")
	}
}