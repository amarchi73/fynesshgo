package types

import "fyne.io/fyne/v2/data/binding"

type Dati struct {
	Nome     binding.String
	User     binding.String
	Host     binding.String
	Pass     binding.String
	Xwindows binding.Bool
}

//{"nome":"adler","user":"adler","host":"nomehost","pass":"password"}
type DatiJson struct {
	Nome     string `json:"nome"`
	User     string `json:"user"`
	Host     string `json:"host"`
	Pass     string `json:"pass"`
	Xwindows bool   `json:"xwindows"`
}

type ElencoDatiJson struct {
	Nome string `json:"nome"`
	Path string `json:"path"`
}

var Data []DatiJson
var ElencoData []ElencoDatiJson

var ConfigPath string
var DataPath string

const ConfigNonValido = -2

var FileJSON = "data.json"
var DefaultFileJson = ""
