package utils

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

var comando string = "/usr/bin/pass"
var passPath string = "fynesshgo"
var recurs int = 0

func SalvaElencoPass(nome string, valore string) {
	cmd := exec.Command(comando, "insert", passPath+"/"+nome, "-f", "-m")
	//cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("\n" + valore)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	/*stdout, err := cmd.Output()
	fmt.Println(stdout, err)
	fmt.Println("======================")*/
}
func LeggiElencoPass(nome string) string {
	/*
	** Sempre meglio mettere l'elefante al Cairo
	** https://github.com/amarchi73/mysshgui/blob/master/elefante.md
	 */
	if recurs > 4 {
		return ""
	}
	cmd := exec.Command(comando, passPath+"/"+nome)
	stdout, err := cmd.Output()
	if err != nil {
		SalvaElencoPass(nome, "\n")
		recurs++
		return LeggiElencoPass(nome)
	}
	return string(stdout)
}
