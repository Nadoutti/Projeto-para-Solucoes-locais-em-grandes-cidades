package services

import (
	"log"
	"projSolucoesLocais/models"
	"projSolucoesLocais/repository"
)

func CreateReport (createreportdata models.CreateReport) (models.Report, error) {
  report, err := repository.CreateReport(createreportdata) 

  if err != nil {
    log.Println("Erro ao criar o report", err)
    return models.Report{}, err
  }

  return report, nil
}

func FetchAllReports () ([]models.Report, error) {
  reports, erro := repository.FetchAllReports()

  if erro != nil {
    log.Println("Erro ao carregar os reports!", erro)
    return nil, erro
  }
  
  return reports, nil
}
