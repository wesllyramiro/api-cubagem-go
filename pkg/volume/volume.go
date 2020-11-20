package volume

import "math"

type Volume struct {
	Psicotropico   int       `json:"psicotropico"`
	Produtos       []Produto `json:"produtos"`
	VolumeRestante float64
}

func (v *Volume) CalcularVolume(m ModeloVolume) {
	v.VolumeRestante = m.CalcularVolume()
}

func (v *Volume) QunatidadeCabeNoVolume(p Produto) int {
	result := (v.VolumeRestante / p.Volumetria)

	return int(result)
}

func (v *Volume) AdicionarProduto(p Produto) int {
	if !v.CabeNoVolume(p) {

		if v.VolumeEstáVazio() {
			return p.Quantidade
		}else {
			return 0
		}
	}

	if p.VolumetriaTotal() < v.VolumeRestante {
		v.Produtos = append(v.Produtos, p)
	} else {
		qtd := v.QunatidadeCabeNoVolume(p)

		p.AlterarQuantidade(qtd)

		v.Produtos = append(v.Produtos, p)
	}

	v.DiminuirVolumetriaRestante(p.VolumetriaTotal())

	return p.Quantidade
}

func (v *Volume) DiminuirVolumetriaRestante(volumetria float64) {
	v.VolumeRestante -= volumetria
}

func (v *Volume) CabeNoVolume(p Produto) bool {
	
	quantidade := v.QunatidadeCabeNoVolume(p)

	if quantidade <= 0 {
		return false
	}

	return true
}

func (v *Volume) VolumeEstáVazio() bool {
	tamanho := len(v.Produtos)

	if tamanho <= 0 {
		return true
	}else {
		return false
	}
}

type ModeloVolume struct {
	Peso           float64 `json:"peso"`
	Capacidade     float64 `json:"capacidade"`
	Aproveitamento float64 `json:"aproveitamento"`
	VolumeRestante float64 `json:"volumeRestante"`
}

func (m *ModeloVolume) CalcularVolume() float64 {
	return math.Round(m.Capacidade - (m.Capacidade * ((100 - m.Aproveitamento) / 100)))
}
