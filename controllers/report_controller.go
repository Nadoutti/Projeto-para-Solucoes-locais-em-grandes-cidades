package controllers

import (
	"net/http"
	"projSolucoesLocais/models"
	"projSolucoesLocais/services"

	"github.com/gin-gonic/gin"
)

func PostReport(c *gin.Context) {
  var createReportData models.CreateReport

  if err := c.ShouldBindJSON(&createReport); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados invalidos"})
    return
  }

  report, err := services.CreateReport(createReportData)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro para criar o report"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"mensagem": "Report criado com sucesso", "report": report})

  
}
