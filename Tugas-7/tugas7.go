package main

import (
	"fmt"
)

// ** SOAL 1 *
type Buah struct {
    Nama        string
    Warna       string
    AdaBijinya  bool
    Harga       int
}

// ** SOAL 2 * 
type segitiga struct{
	alas, tinggi int
  }
  
type persegi struct{
	sisi int
}

type persegiPanjang struct{
	panjang, lebar int
}

func (s segitiga) luas() float64 {
	return 0.5 * float64(s.alas) * float64(s.tinggi)
}

func (p persegi) luas() float64 {
	return float64(p.sisi) * float64(p.sisi)
}

func (pp persegiPanjang) luas() float64 {
	return float64(pp.panjang) * float64(pp.lebar)
}

// ** SOAL 3 *
type phone struct{
	name, brand string
	year int
	colors []string
 }

 func (p *phone) addColor(newColor string) {
	p.colors = append(p.colors, newColor)
 }

 // ** SOAL 4 *
 type movie struct {
	title, genre string
	duration, year int
 }

 func tambahDataFilm(title string,  duration int, genre string, year int, dataFilm *[]movie) {
	newMovie := movie{
		title:    title,
		genre:    genre,
		duration: duration,
		year:     year,
	}

	*dataFilm = append(*dataFilm, newMovie)
 }

func main() {
	// ** SOAL 1 *
	nanas := Buah{
        Nama:       "Nanas",
        Warna:      "Kuning",
        AdaBijinya: false,
        Harga:      9000,
    }

    jeruk := Buah{
        Nama:       "Jeruk",
        Warna:      "Oranye",
        AdaBijinya: true,
        Harga:      8000,
    }

    semangka := Buah{
        Nama:       "Semangka",
        Warna:      "Hijau & Merah",
        AdaBijinya: true,
        Harga:      10000,
    }

    pisang := Buah{
        Nama:       "Pisang",
        Warna:      "Kuning",
        AdaBijinya: false,
        Harga:      5000,
    }

    fmt.Println("Buah Nanas:")
    fmt.Println("Nama:", nanas.Nama)
    fmt.Println("Warna:", nanas.Warna)
    fmt.Println("Ada Bijinya:", nanas.AdaBijinya)
    fmt.Println("Harga:", nanas.Harga)
    fmt.Println()

    fmt.Println("Buah Jeruk:")
    fmt.Println("Nama:", jeruk.Nama)
    fmt.Println("Warna:", jeruk.Warna)
    fmt.Println("Ada Bijinya:", jeruk.AdaBijinya)
    fmt.Println("Harga:", jeruk.Harga)
    fmt.Println()

    fmt.Println("Buah Semangka:")
    fmt.Println("Nama:", semangka.Nama)
    fmt.Println("Warna:", semangka.Warna)
    fmt.Println("Ada Bijinya:", semangka.AdaBijinya)
    fmt.Println("Harga:", semangka.Harga)
    fmt.Println()

    fmt.Println("Buah Pisang:")
    fmt.Println("Nama:", pisang.Nama)
    fmt.Println("Warna:", pisang.Warna)
    fmt.Println("Ada Bijinya:", pisang.AdaBijinya)
    fmt.Println("Harga:", pisang.Harga)


	// ** SOAL 2 *
	segitiga := segitiga{alas: 5, tinggi: 8}
	persegi := persegi{sisi: 4}
	persegiPanjang := persegiPanjang{panjang: 6, lebar: 3}

	fmt.Println("Luas Segitiga:", segitiga.luas())
	fmt.Println("Luas Persegi:", persegi.luas())
	fmt.Println("Luas Persegi Panjang:", persegiPanjang.luas())

	// ** SOAL 3 *
	iphone := phone{
		name:   "Iphone 16",
		brand:  "Apple",
		year:   2024,
		colors: []string{"White", "Black"},
	}

	fmt.Println("Warna sebelum penambahan:", iphone.colors)

	iphone.addColor("Blue")
	iphone.addColor("Gold")

	fmt.Println("Warna setelah penambahan:", iphone.colors)

	// ** SOAL 4 *
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. title: %s\nduration: %d jam\ngenre: %s\nyear: %d\n\n", i+1, film.title, film.duration, film.genre, film.year)
	}
}