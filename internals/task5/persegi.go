package task5

type Balok struct {
	Sisi float64
}

func SetSisiBalok(sisi float64) *Balok {
	return &Balok{
		Sisi: sisi,
	}
}

func (k *Balok) SetSisi(sisi float64) {
	k.Sisi = sisi
}

func (k *Balok) Luas() float64 {
	return 6 * (k.Sisi * k.Sisi)
}

func (k *Balok) Keliling() float64 {
	return 12 * k.Sisi
}

func (k *Balok) Volume() float64 {
	return k.Sisi * k.Sisi * k.Sisi
}
