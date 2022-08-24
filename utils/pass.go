package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var ComandoPass string = "/usr/bin/pass"
var PassPath string = "fynesshgo"
var recurs int = 0

func SalvaElencoPass(nome string, valore string) {
	nome = strings.ReplaceAll(nome, ".json", "")

	cmd := exec.Command(ComandoPass, "insert", PassPath+"/"+nome, "-f", "-m")
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
	** https://github.com/amarchi73/fynesshgo/blob/master/elefante.md
	 */
	if recurs > 4 {
		return ""
	}
	nome = strings.ReplaceAll(nome, ".json", "")
	fmt.Println(ComandoPass, PassPath+"/"+nome)

	cmd := exec.Command(ComandoPass, PassPath+"/"+nome)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("...", PassPath+"/"+nome, stdout)
		if strings.Contains(err.Error(), "exit status 2") {
			return "NO"
		}
		fmt.Println("-", err)
		SalvaElencoPass(nome, "\n")
		recurs++
		return LeggiElencoPass(nome)
	}
	return string(stdout)
}
