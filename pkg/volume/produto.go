package volume

type Produto struct {
	Codigo        int    `json:"codigo"`
	Endereco      string `json:"endereco"`
	Quantidade    int    `json:"quantidade"`
	EstoqueMinimo int    `json:"estoqueMinimo"`
	Volumetria	  float64 
}

func (p *Produto) AlterarQuantidade(qtd int) {
	p.Quantidade = qtd
}
func (p *Produto) VolumetriaTotal() float64 {
	return p.Volumetria * float64(p.Quantidade)
}

type ProdutoDetalhe struct {
	Codigo           int
	EstoqueMinino    int
	QtdProduto       int
	EnderecoDeposito string
	VolumeProduto    float64
	EnderecoReduzido string
}

func (p *ProdutoDetalhe) DiminuirQtdJaAddNoVolume(qtd int) {
	p.QtdProduto -= qtd
}

func (p *ProdutoDetalhe) AindaTemQuantidade() bool {
	if p.QtdProduto > 0 {
		return true
	}else {
		return false
	} 
}

func (p *ProdutoDetalhe) TotalVolume() float64 {
	return p.VolumeProduto * float64(p.QtdProduto)
}
