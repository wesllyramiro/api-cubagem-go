package volume

type IService interface {
	BuscarModeloVolume(int) (ModeloVolume, error)
	RealizarCubagem(int, int, bool) ([]Volume, error)
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
func (s *service) RealizarCubagem(filial int, idModelo int, isPsico bool) ([]Volume, error) {

	modelo, err := s.r.BuscarModeloVolume(idModelo)
	if err != nil {
		return nil, err
	}

	enderecos, err := s.r.BuscarProdutosCubagem(filial, isPsico)
	if err != nil {
		return nil, err
	}

	var vs []Volume

	var v Volume
	v.CalcularVolume(modelo)

	for _, endereco := range enderecos {
		for endereco.AindaTemProduto() {

			vl := AdicionarProdutoAoVolume(&endereco, modelo, &vs, &v)

			if vl != nil {
				vs = append(vs, v)

				v = Volume{}
				v.CalcularVolume(modelo)
			}
		}
	}

	return vs, err
}

func AdicionarProdutoAoVolume(e *Endereco, m ModeloVolume, vs *[]Volume, v *Volume) *Volume {

	for _, pr := range e.ProdutosComQuntidade() {
		p := Produto{
			Codigo:        pr.Codigo,
			Endereco:      pr.EnderecoDeposito,
			EstoqueMinimo: pr.EstoqueMinino,
			Quantidade:    pr.QtdProduto,
			Volumetria:    pr.VolumeProduto	}

		qtdAdicionado := v.AdicionarProduto(p)

		e.DiminuirQtdJaAddNoVolume(pr, qtdAdicionado)
	}

	if e.AindaTemProduto() {
		return v
	}else {
		return nil
	}
}
