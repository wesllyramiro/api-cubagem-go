package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/wesllyramiro/api-cubagem-go/config"
	"github.com/wesllyramiro/api-cubagem-go/pkg/http/rest"
	"github.com/wesllyramiro/api-cubagem-go/pkg/volume"
)

func main() {

	dbSource := config.GetStringConn("api-cubagem-go")

	con, err := sql.Open("sqlserver", dbSource)
	if err != nil {
		log.Fatal("cannot connection sqlserver:", err)
	}
	//Repositories
	volumeRepo := volume.NewRepo(con)
	//Services
	volumeService := volume.NewService(volumeRepo)

	router := start(volumeService)

	err = router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func start(v volume.IService) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("cubagem/encaixotar-produtos", rest.EncaixotarProdutos(v))
	}

	return router
}
