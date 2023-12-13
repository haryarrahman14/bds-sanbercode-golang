package main

import "fmt"

func main() {

	// ** SOAL 1 *
	for i := 1; i <= 20; i++ {
        if i%2 == 1 {
            fmt.Println(i, "- Santai")
        } else {
            fmt.Println(i, "- Berkualitas")
        }
        if i%3 == 0 && i%2 == 1 {
            fmt.Println(i, "- I Love Coding")
        }
    }

	// ** SOAL 2 *
	for i := 1; i <= 7; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("#")
        }
        fmt.Println()
    }

	// ** SOAL 3 *
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
    hasil := kalimat[2:7]
    fmt.Println(hasil)

	// ** SOAL 4 *
	var sayuran = []string{}
    sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
    for i, s := range sayuran {
        fmt.Printf("%d. %s\n", i+1, s)
    }

	// ** SOAL 5 *
	var satuan = map[string]int{
        "panjang": 7,
        "lebar":   4,
        "tinggi":  6,
    }
    for k, v := range satuan {
        fmt.Printf("%s = %d\n", k, v)
    }
    volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
    fmt.Printf("Volume Balok = %d\n", volume)
}