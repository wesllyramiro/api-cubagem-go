package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wesllyramiro/api-cubagem-go/pkg/volume"
)

func GetModeloVolume(s volume.IService) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Query("id"))

		vol, err := s.BuscarModeloVolume(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, responseMessage("Modelo não encontrado"))
			return
		}

		ctx.JSON(http.StatusOK, vol)
	}
}

type encaixotarProdutosRequest struct {
	IdModelo string `form:"idModelo" binding:"required"`
	Filial   string `form:"filial" binding:"required"`
}

func EncaixotarProdutos(s volume.IService) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Query("IdModelo"))
		filial, _ := strconv.Atoi(ctx.Query("Filial"))
		
		psico := [2]bool { true, false }

		var vs []volume.Volume

		for _, isPsico := range psico {
			vls, err := s.RealizarCubagem(filial, id, isPsico)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, err) 
			}

			vs = append(vs, vls...)
		}
		
		ctx.JSON(http.StatusOK, vs)
	}
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
func responseMessage(err string) gin.H {
	return gin.H{"message": err}
}
