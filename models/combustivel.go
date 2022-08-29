package models

type Combustivel struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

var Combustiveis []Combustivel
