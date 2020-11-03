package volume

type Produto struct {
	Codigo        int    `json:"codigo"`
	Endereco      string `json:"endereco"`
	Quantidade    int    `json:"quantidade"`
	EstoqueMinimo int    `json:"estoqueMinimo"`
}
type ProdutoDetalhe struct {
	Codigo           int
	Filial           int
	DescricaoFilial  string
	Digito           int
	Descricao        string
	TipoProduto      string
	VlVenda          float64
	EstoqueMinino    int
	QuantidadeEnviar int
	Barras           string
	QtdProduto       int
	EnderecoDeposito string
	VolumeProduto    float64
	EnderecoReduzido string
}
