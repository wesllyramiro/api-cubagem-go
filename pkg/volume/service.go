package volume

type IService interface {
	BuscarModeloVolume(int) (ModeloVolume, error)
	RealizarCubagem(int, int) []Endereco
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
func (s *service) RealizarCubagem(filial int, idModelo int) []Endereco {
	end, _ := s.r.BuscarProdutosCubagem(filial, false)

	return end
}
