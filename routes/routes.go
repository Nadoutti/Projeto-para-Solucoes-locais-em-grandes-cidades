package routes

import (
	"projSolucoesLocais/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
  router := gin.Default()
  

  report := router.Group("/report")
  {
    report.POST("", controllers.PostReport)
  }

  return router
}
