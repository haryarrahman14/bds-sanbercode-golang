package model

import "math"

type Tabung struct{
	JariJari, Tinggi float64
}

type Balok struct{
	Panjang, Lebar, Tinggi float64
}

type HitungBangunRuang interface{
	Volume()		float64
	LuasPermukaan() float64
}


func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * float64(t.Tinggi)
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (float64(t.JariJari) + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok)  LuasPermukaan() float64 {
	return 2 * float64((b.Panjang * b.Lebar) + (b.Panjang * b.Tinggi) + (b.Lebar * b.Tinggi))
}
