package models

type Posto struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

var Postos []Posto

/*
func BuscaPosto() {
	db := database.ConectDB()
	defer db.Close()

	selectPosto, err := db.Query("SELECT * FROM posto_combustivel ORDER BY id ASC")

	if err != nil {
		panic(err.Error())
	}

	p := Posto{}

	postos := []Posto{}

	for selectPosto.Next() {
		var id int
		var nome string

		err = selectPosto.Scan(&id, &nome)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome

		postos = append(postos, p)
	}

	return postos

}
*/
