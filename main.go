package main

import (
	"projSolucoesLocais/routes"

	"github.com/gin-gonic/gin"
)

func main()  {
  var db = "meu db"
  r := gin.Default() 
  routes.SetupRoutes(r, db)
}


