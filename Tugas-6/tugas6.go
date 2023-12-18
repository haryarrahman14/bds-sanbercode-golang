package main

import (
	"fmt"
	"math"
)

// ** SOAL 1 **
func hitungLingkaran(jariJari float64, luasLingkaran *float64, kelilingLingkaran *float64) {
	*luasLingkaran = math.Pi * math.Pow(jariJari, 2)
	*kelilingLingkaran = 2 * math.Pi * (jariJari)
}

// ** SOAL 2 **
func introduce(sentence *string, name string, gender string, job string, age string) {
	*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, job, age)
	if gender == "perempuan" {
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", name, job, age)
	}
}

// ** SOAL 3 **
func tambahBuah(buah *[]string, namaBuah string) {
	*buah = append(*buah, namaBuah)
}

// ** SOAL 4 **
func tambahDataFilm( title string, duration string, genre string, year string, film *[]map[string]string,) {
	*film = append(*film, map[string]string{
		"title": title,
		"duration": duration,
		"genre": genre,
		"year": year,
	})
}

func main() {
	// ** SOAL 1 **
	var luasLingkaran float64 
	var kelilingLingkaran float64

	var jariJariLingkaran float64 = 7
	hitungLingkaran(jariJariLingkaran, &luasLingkaran, &kelilingLingkaran)

	fmt.Println("Luas Lingkaran:", luasLingkaran)
	fmt.Println("Keliling Lingkaran:", kelilingLingkaran)

	// ** SOAL 2 **
	var sentence string 
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence)

	// ** SOAL 3 **
	var buah = []string{}

	tambahBuah(&buah, "Jeruk")
	tambahBuah(&buah, "Semangka")
	tambahBuah(&buah, "Mangga")
	tambahBuah(&buah, "Strawberry")
	tambahBuah(&buah, "Durian")
	tambahBuah(&buah, "Manggis")
	tambahBuah(&buah, "Alpukat")

	fmt.Println("Daftar Buah:")
	for i, namaBuah := range buah {
		fmt.Printf("%d. %s\n", i+1, namaBuah)
	}

	// ** SOAL 4 **
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. title: %s\n", i+1, film["title"])
		fmt.Printf("   duration: %s\n", film["duration"])
		fmt.Printf("   genre: %s\n", film["genre"])
		fmt.Printf("   year: %s\n", film["year"])
	}
}