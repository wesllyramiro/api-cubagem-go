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

func (v *Volume) QunatidadeCabeNoVolume(p ProdutoDetalhe) int {
	result := (v.VolumeRestante / p.VolumeProduto)

	return int(result)
}

func (v *Volume) ProdutoCabeNoVolume(p ProdutoDetalhe) {
	qtd := v.QunatidadeCabeNoVolume(p)

	if qtd == p.QtdProduto {

	}
}

func (v *Volume) AdicionarProduto(p Produto, volumeDoProduto float64, qtd int) {
	p.AdicionarQuantidade(qtd)

	v.VolumeRestante -= float64(p.Quantidade) * volumeDoProduto
	v.Produtos = append(v.Produtos, p)
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
