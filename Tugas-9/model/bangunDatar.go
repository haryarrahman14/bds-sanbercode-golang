package model

type SegitigaSamaSisi struct{
	Alas, Tinggi float64
}

type PersegiPanjang struct{
	Panjang, Lebar float64
}

type HitungBangunDatar interface{
	Luas() 		float64
	Keliling()	float64
}

func (s SegitigaSamaSisi) Luas() float64 {
	return float64(s.Alas * s.Alas) / 2
}

func (s SegitigaSamaSisi) Keliling() float64 {
	return 3 * float64(s.Alas)
}

func (pp PersegiPanjang) Luas() float64 {
	return float64(pp.Panjang * pp.Lebar)
}

func (pp PersegiPanjang) Keliling() float64 {
	return 2 * float64(pp.Panjang + pp.Lebar)
}