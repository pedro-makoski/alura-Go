package routes

import (
	"pessoas/controllers"

	"github.com/gin-gonic/gin"
)

func GenRoutes() {
	r := gin.Default()

	r.GET(controllers.ROUTES_BOAS_VINDAS, controllers.BoasVindas)
	r.GET(controllers.MOSTRAR_ID_ROUTE, controllers.MostrarId)

	r.POST(controllers.ROUTE_POST_PESSOA, controllers.PostPessoa)

	r.Run(":8000")
}
