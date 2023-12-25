package helper

import "math"

func HitungLuasLingkaran(jariJari float64) float64 {
	luas := math.Pi * math.Pow(jariJari, 2)
	return math.Round(luas)
}

func HitungKelilingLingkaran(jariJari float64) float64 {
	keliling := 2 * math.Pi * jariJari
	return math.Round(keliling)
}
