package models


type CreateReport struct {
  Titulo string `json:"titulo"`
  Descricao string `json:"descrição"`
  Gravidade int `json:"gravidade"`
}

type Report struct {
  Id string `json:"id"`
  Titulo string `json:"titulo"`
  Descricao string `json:"descrição"`
  Gravidade int `json:"gravidade"`
  Estado string `json:"estado"`
}


