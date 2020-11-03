package volume

import (
	"database/sql"
)

type IRepository interface {
	BuscarModeloVolume(int) (ModeloVolume, error)
	BuscarProdutosCubagem(int, bool) ([]Endereco, error)
}

type repository struct {
	db *sql.DB
}

func NewRepo(d *sql.DB) IRepository {
	return &repository{d}
}

const buscarModelo = `
SELECT  
	DLCR_PESO 							as Peso,           
	DLCR_CAPACIDADE 					as Capacidade,
	DLCR_APROVEITAMENTO 				as Aproveitamento		
FROM COSMOSPDP..DEVOLUCAO_LOJA_CAIXA_REVERSA
	WHERE	
	DLCR_ID_CAIXA = @p1
`

func (r *repository) BuscarModeloVolume(id int) (ModeloVolume, error) {
	row := r.db.QueryRow(buscarModelo, id)
	var i ModeloVolume
	err := row.Scan(
		&i.Peso,
		&i.Capacidade,
		&i.Aproveitamento)
	return i, err
}

const BuscarDetalheProduto = `		
		SELECT 
				SUBSTRING(EnderecoDeposito, 1,7)	as EnderecoReduzido
				,Codigo								as Produtos_Codigo
				,Filial								as Produtos_Filial					
				,DescricaoFilial					as Produtos_DescricaoFilial				
				,Digito								as Produtos_Digito
				,Descricao							as Produtos_Descricao
				,TipoProduto						as Produtos_TipoProduto
				,VlVenda							as Produtos_VlVenda
				,EstoqueMinino						as Produtos_EstoqueMinino
				,QuantidadeEnviar					as Produtos_QuantidadeEnviar
				,Barras								as Produtos_Barras				
				,QtdProduto							as Produtos_QtdProduto
				,EnderecoDeposito					as Produtos_EnderecoDeposito				
				,VolumeProduto						as Produtos_VolumeProduto
		FROM 
				VW_OBTER_PRODUTOS_REVERSA
		WHERE 
				filial = @p1			
				AND	(ALTURA IS NOT null	or				
				Comprimento IS NOT null	or			
				LARGURA	IS NOT null)				
`
const ExpurgaPsicotropico = `
	AND TipoProduto <> 'Psicotrópico'
`
const PorPsicotropicos = `
	AND TipoProduto = 'Psicotrópico'
`
const OrdenaProdutos = `
	ORDER BY TipoProduto, SUBSTRING(EnderecoDeposito, 1, 7), VolumeProduto desc
`

func (r *repository) BuscarProdutosCubagem(filial int, isPsico bool) ([]Endereco, error) {

	var condition string
	if isPsico {
		condition = PorPsicotropicos
	} else {
		condition = ExpurgaPsicotropico
	}

	var enderecosMap = make(map[string][]ProdutoDetalhe)

	rows, err := r.db.Query(BuscarDetalheProduto+condition+OrdenaProdutos, filial)
	defer rows.Close()

	for rows.Next() {
		var p ProdutoDetalhe
		rows.Scan(
			&p.EnderecoReduzido,
			&p.Codigo,
			&p.Filial,
			&p.DescricaoFilial,
			&p.Digito,
			&p.Descricao,
			&p.TipoProduto,
			&p.VlVenda,
			&p.EstoqueMinino,
			&p.QuantidadeEnviar,
			&p.Barras,
			&p.QtdProduto,
			&p.EnderecoDeposito,
			&p.VolumeProduto)

		if produtos, ok := enderecosMap[p.EnderecoReduzido]; ok {
			enderecosMap[p.EnderecoReduzido] = append(produtos, p)
		} else {
			enderecosMap[p.EnderecoReduzido] = []ProdutoDetalhe{p}
		}
	}
	enderecos := []Endereco{}
	for key, val := range enderecosMap {
		enderecos = append(enderecos, Endereco{
			EnderecoReduzido: key,
			Produtos:         val})
	}

	return enderecos, err
}
