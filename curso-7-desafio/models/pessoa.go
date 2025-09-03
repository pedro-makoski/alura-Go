package models

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Ativo bool   `json:"ativo"`
}

var Pessoas = []Pessoa{
	{
		Nome:  "Joana",
		Idade: 10,
		Ativo: false,
	},
	{
		Nome:  "Joana 2",
		Idade: 15,
		Ativo: false,
	},
	{
		Nome:  "Joana 3",
		Idade: 30,
		Ativo: false,
	},
	{
		Nome:  "Joana 10",
		Idade: 60,
		Ativo: true,
	},
}
