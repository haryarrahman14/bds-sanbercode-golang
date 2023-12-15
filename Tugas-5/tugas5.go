package main

import (
	"fmt"
	"strings"
)

// ** SOAL 1 **
func luasPersegipanjang(p, l int) int {
	return p * l
}

func kelilingPersegipanjang(p, l int) int {
	return 2 * (p + l)
}

func volumeBalok(p, l, t int) int {
	return p * l * t
}

// ** SOAL 2 **
func introduce(name string, gender string, job string, age string) (result string) {
	var prefix string
	titleWord := "adalah seorang"
	ageWord := fmt.Sprintf("yang berusia %s tahun", age)

	if (gender == "laki-laki") {
		prefix = "Pak"
	} else if (gender == "perempuan") {
		prefix = "Bu"
	}

	result = fmt.Sprintf("%s %s %s %s %s", prefix, name, titleWord, job, ageWord)

	return
}

// ** SOAL 3 **
func buahFavorit(name string, buah ...string) string {
	prefix := "halo nama saya"
	prefixBuahFavorit := "dan buah favorit saya adalah"
	buahFavoritStr := strings.Join(buah, "\", \"")

	result := fmt.Sprintf("%s %s %s \"%s\".", prefix, name, prefixBuahFavorit, buahFavoritStr)
	return result
}

func main() {
	// ** SOAL 1 **
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegipanjang(panjang, lebar)
	keliling := kelilingPersegipanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// ** SOAL 2 **
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) 

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) 

	// ** SOAL 3 **
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	// ** SOAL 4 **
	var dataFilm = []map[string]string{}

	var tambahDataFilm = func(title string, duration string, genre string, year string) {
		var newFilm = map[string]string{}
		newFilm["title"] = title
		newFilm["jam"] = duration
		newFilm["genre"] = genre
		newFilm["tahun"] = year
		dataFilm = append(dataFilm, newFilm)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}


}