package volume

import "math"

type Volume struct {
	Psicotropico int       `json:"psicotropico"`
	Produtos     []Produto `json:"produtos"`
}

type ModeloVolume struct {
	Peso           float64 `json:"peso"`
	Capacidade     float64 `json:"capacidade"`
	Aproveitamento float64 `json:"aproveitamento"`
	VolumeRestante float64 `json:"volumeRestante"`
}

func (v *ModeloVolume) CalcularVolume() float64 {
	return math.Round(v.Capacidade - (v.Capacidade * ((100 - v.Aproveitamento) / 100)))
}
