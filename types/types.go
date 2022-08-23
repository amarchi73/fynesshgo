package types

import "fyne.io/fyne/v2/data/binding"

type Dati struct {
	Nome binding.String
	User binding.String
	Host binding.String
	Pass binding.String
}

//{"nome":"adler","user":"adler","host":"nomehost","pass":"password"}
type DatiJson struct {
	Nome string `json:"nome"`
	User string `json:"user"`
	Host string `json:"host"`
	Pass string `json:"pass"`
}

var Data []DatiJson

const FileJSON = "data.json"
