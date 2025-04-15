package routes

import (
	"projSolucoesLocais/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
  router := gin.Default()

  publicRoutes := map[string]bool{
    "POST, /api/login": true,
    "POST, /api/register": true,
    "POST, /api/forgot_password": true,
    "POST, /api/validar_existencia": true,

  }
  

  report := router.Group("/report")
  {
    report.POST("", controllers.PostReport)
    report.GET("", controllers.GetAllReports)
  }

  return router
}
