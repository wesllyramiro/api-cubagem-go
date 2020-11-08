package volume

type IService interface {
	BuscarModeloVolume(int) (ModeloVolume, error)
	RealizarCubagem(int, int) ([]Endereco, error)
}

type service struct {
	r IRepository
}

func NewService(rp IRepository) IService {
	return &service{rp}
}

func (s *service) BuscarModeloVolume(id int) (ModeloVolume, error) {
	vol, err := s.r.BuscarModeloVolume(id)

	return vol, err
}
func (s *service) RealizarCubagem(filial int, idModelo int) ([]Endereco, error) {

	modelo, err := s.r.BuscarModeloVolume(idModelo)
	if err != nil {
		return nil, err
	}

	enderecos, err := s.r.BuscarProdutosCubagem(filial, false)
	if err != nil {
		return nil, err
	}

	for _, endereco := range enderecos {
		AdicionarProdutoAoVolume(endereco, modelo)
	}

	return enderecos, err
}

func AdicionarProdutoAoVolume(e Endereco, m ModeloVolume) {
	var v Volume
	v.CalcularVolume(m)

	for _, produto := range e.Produtos {
		p := Produto{
			Codigo:        produto.Codigo,
			Endereco:      produto.EnderecoDeposito,
			EstoqueMinimo: produto.EstoqueMinino}

		qtd := v.QunatidadeCabeNoVolume(produto)

		if qtd < produto.QtdProduto {
			p.AdicionarQuantidade(qtd)
			v.AdicionarProduto(p, produto.VolumeProduto)
		} else {
			p.AdicionarQuantidade(produto.QtdProduto)
			v.AdicionarProduto(p, produto.VolumeProduto)
		}
	}
}
