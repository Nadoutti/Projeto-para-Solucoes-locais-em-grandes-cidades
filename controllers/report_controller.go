package controllers

import (
	"net/http"
	"projSolucoesLocais/models"
	"projSolucoesLocais/services"

	"github.com/gin-gonic/gin"
)

func PostReport(c *gin.Context) {
  var createReportData models.CreateReport

  if err := c.ShouldBindJSON(&createReportData); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados invalidos"})
    return //parar de rodar a funcao se der merda
  }

  report, err := services.CreateReport(createReportData)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro para criar o report"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"mensagem": "Report criado com sucesso", "report": report})

  
}

func GetAllReports(c *gin.Context) {
  reports, err := services.FetchAllReports()
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"erro": "Nao foi encontrado report algum"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"mensagem": "Reports listados!", "reports": reports})
}
