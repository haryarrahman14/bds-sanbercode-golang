package main

import (
	"bds-sanbercode-golang/Tugas-10/helper"
	"flag"
	"fmt"
	"sort"
	"time"
)



func main() {
	kalimat := "Golang Backend Development"			// ** SOAL 1 *
	tahun := 2021									// ** SOAL 1 *

	angka := 1										// ** SOAL 3 *

	defer helper.TampilkanKalimatDanTahun(kalimat, tahun)	// ** SOAL 1 *
	defer helper.CetakAngka(&angka)						// ** SOAL 3 *
	

	// ** SOAL 2 *
	fmt.Println(helper.KelilingSegitigaSamaSisi(4, true))
	fmt.Println(helper.KelilingSegitigaSamaSisi(8, false))
	fmt.Println(helper.KelilingSegitigaSamaSisi(0, true))
	fmt.Println(helper.KelilingSegitigaSamaSisi(0, false))

	// ** SOAL 3 *
	helper.TambahAngka(7, &angka)
	helper.TambahAngka(6, &angka)
	helper.TambahAngka(-1, &angka)
	helper.TambahAngka(9, &angka)

	// ** SOAL 4 *
	var phones = []string{}

	helper.TambahData("Xiaomi", &phones)
	helper.TambahData("Asus", &phones)
	helper.TambahData("IPhone", &phones)
	helper.TambahData("Samsung", &phones)
	helper.TambahData("Oppo", &phones)
	helper.TambahData("Realme", &phones)
	helper.TambahData("Vivo", &phones)

	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second)
	}

	// ** SOAL 5 *
	jariJari1 := 7.0
	luas1 := helper.HitungLuasLingkaran(jariJari1)
	keliling1 := helper.HitungKelilingLingkaran(jariJari1)
	fmt.Printf("Jari-jari: %.f, Luas: %.f, Keliling: %.f\n", jariJari1, luas1, keliling1)

	jariJari2 := 10.0
	luas2 := helper.HitungLuasLingkaran(jariJari2)
	keliling2 := helper.HitungKelilingLingkaran(jariJari2)
	fmt.Printf("Jari-jari: %.f, Luas: %.f, Keliling: %.f\n", jariJari2, luas2, keliling2)

	jariJari3 := 15.0
	luas3 := helper.HitungLuasLingkaran(jariJari3)
	keliling3 := helper.HitungKelilingLingkaran(jariJari3)
	fmt.Printf("Jari-jari: %.f, Luas: %.f, Keliling: %.f\n", jariJari3, luas3, keliling3)

	// ** SOAL 6 *
	var panjang, lebar float64

	flag.Float64Var(&panjang, "panjang", 7, "Panjang persegi panjang")
	flag.Float64Var(&lebar, "lebar", 8, "Lebar persegi panjang")
	flag.Parse()

	luas := helper.HitungLuasPersegiPanjang(panjang, lebar)
	keliling := helper.HitungKelilingPersegiPanjang(panjang, lebar)

	fmt.Printf("Luas Persegi Panjang: %.2f\n", luas)
	fmt.Printf("Keliling Persegi Panjang: %.2f\n", keliling)
}

