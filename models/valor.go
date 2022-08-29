package models

type Valor struct {
	Id             int     `json:"id"`
	Valor          float64 `json:"valor"`
	Id_combustivel int     `json:"id_combustivel"`
	Id_posto       int     `json:"id_posto"`
}

var Valores []Valor
