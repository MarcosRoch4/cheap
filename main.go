package main

import (
	"fmt"

	"github.com/MarcosRoch4/cheap/routes"
)

func main() {
	/*
		models.Combustiveis = []models.Combustivel{
			{Id: 1, Nome: "Posto 1"},
		}
	*/
	fmt.Println("Iniciando o Servidor Rest em GO")

	routes.HandleRequest()
}
