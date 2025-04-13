package task5

type Kubus struct {
	Sisi float64
}

func SetSisi(sisi float64) *Kubus {
	return &Kubus{
		Sisi: sisi,
	}
}

func (k *Kubus) SetSisi(sisi float64) {
	k.Sisi = sisi
}

func (k *Kubus) Luas() float64 {
	return 6 * (k.Sisi * k.Sisi)
}

func (k *Kubus) Keliling() float64 {
	return 12 * k.Sisi
}

func (k *Kubus) Volume() float64 {
	return k.Sisi * k.Sisi * k.Sisi
}
