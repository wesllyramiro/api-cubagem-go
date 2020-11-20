package volume

type Endereco struct {
	EnderecoReduzido string
	Produtos         []ProdutoDetalhe
}

func (e *Endereco) DiminuirQtdJaAddNoVolume(p ProdutoDetalhe, qtd int) {
	pn := []ProdutoDetalhe{}
	for _, pr := range e.Produtos {
		if pr == p {
			pr.DiminuirQtdJaAddNoVolume(qtd)
		}
		pn = append(pn, pr)
	}  
	e.Produtos = pn
}
func (e *Endereco)  ProdutosComQuntidade() []ProdutoDetalhe {
	p := []ProdutoDetalhe{}

	for _, pr := range e.Produtos {
		if pr.AindaTemQuantidade() {
            p = append(p, pr)
        }
	}

	e.Produtos = p
	return e.Produtos
}

func (e *Endereco) AindaTemProduto() bool {
	p := e.ProdutosComQuntidade()

	tamanho := len(p)

	if tamanho > 0 {
		return true
	}else {
		return false
	}
}