package services

import (
	"log"
	"projSolucoesLocais/models"
)

func CreateReport (createreportdata models.CreateReport) (models.Report, error) {
  report, err := repository.asdjfh(createreportdata)

  if err != nil {
    log.Println("Erro ao criar o report", err)
    return models.Report{}, err
  }

  return report, nil
}
