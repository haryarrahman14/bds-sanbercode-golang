package main

import (
	"fmt"
	"math"
	"strings"
)

// Definisi interface hitung2d
type hitung2d interface {
	luas() float64
	keliling() float64
}

// Definisi interface hitung3d
type hitung3d interface {
	volume() float64
}

// Definisi interface hitung yang menggabungkan hitung2d dan hitung3d
type hitung interface {
	hitung2d
	hitung3d
}

// Struct lingkaran dengan method yang sesuai dengan hitung2d
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// Struct persegi dengan method yang sesuai dengan hitung2d
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// Struct kubus dengan method yang sesuai dengan hitung
type kubus struct {
	sisi float64
}

func (k kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k kubus) keliling() float64 {
	return k.sisi * 12
}

// Fungsi untuk menambahkan data film ke slice dataFilm
type movie struct {
	title    string
	duration int
	genre    string
	year     int
}

var dataFilm = []movie{}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	filmBaru := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*dataFilm = append(*dataFilm, filmBaru)
}

type person struct {
	name 	string
	age 	int
}

func note() {
	// Contoh penerapan interface untuk bangun datar
	var bangunDatar hitung2d

	// Contoh penerapan interface untuk lingkaran
	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())
	fmt.Println()

	// Contoh penerapan interface untuk persegi
	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println()

	// Contoh penerapan interface untuk kubus
	var bangunRuang hitung = kubus{4}
	fmt.Println("===== kubus")
	fmt.Println("luas      :", bangunRuang.luas())
	fmt.Println("keliling  :", bangunRuang.keliling())
	fmt.Println("volume    :", bangunRuang.(kubus).volume())
	fmt.Println()

	// Contoh penggunaan interface kosong
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")
	fmt.Println()

	// Contoh casting variabel interface kosong ke object pointer
	var secretObject interface{} = &person{name: "wick", age: 27}
	var name = secretObject.(*person).name
	fmt.Println(name)
	fmt.Println()

	// Menambahkan data film menggunakan fungsi tambahDataFilm
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	// Menampilkan data film
	fmt.Println("Data Film:")
	for i, film := range dataFilm {
		fmt.Printf("%d. title: %s\n   duration: %d jam\n   genre: %s\n   year: %d\n", i+1, film.title, film.duration, film.genre, film.year)
	}
}
