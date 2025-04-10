package models


type CreateReport struct {
  titulo string `json:"titulo"`
  descricao string `json:"descrição"`
  gravidade int `json:"gravidade"`
}

type Report struct {
  ID string `json:"id"`
  titulo string `json:"titulo"`
  descricao string `json:"descrição"`
  gravidade int `json:"gravidade"`
}


