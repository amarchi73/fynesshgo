package utils

import (
	"bytes"
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
	//fmt.Println(ComandoPass, PassPath+"/"+nome)

	cmd := exec.Command(ComandoPass, PassPath+"/"+nome)
	stdout, err := cmd.Output()
	if err != nil {
		//fmt.Println("...", PassPath+"/"+nome, stdout)
		if strings.Contains(err.Error(), "exit status 2") {
			return "NO"
		}
		//fmt.Println("-", err)
		SalvaElencoPass(nome, "\n")
		recurs++
		return LeggiElencoPass(nome)
	}
	return string(stdout)
}

/*
** -ls			mostra tutti
** -le			mostra elenchi
** -lc=<nome>	mostra le connessioni dell'elenco <nome>
 */
func ListPass(par string) {
	elenchi := GetPass("fynesshgo")
	if strings.Contains(par, "-le") {
		for i := 0; i < len(elenchi); i++ {
			if elenchi[i] == "" {
				continue
			}
			EchoString(elenchi[i])
		}
	} else if strings.Contains(par, "-ls") {
		for i := 0; i < len(elenchi); i++ {
			if elenchi[i] == "" {
				continue
			}
			conn := LeggiElencoDati(elenchi[i])
			for k := 0; k < len(conn); k++ {
				if k == 0 {
					EchoString(elenchi[i])
				}
				if conn[k].Nome == "" {
					continue
				}
				EchoString(" - " + conn[k].Nome)
			}
		}
	} else {
		nome := strings.ReplaceAll(par, "-lc=", "")
		EchoString(nome)
		elenchi := LeggiElencoDati(nome)
		for i := 0; i < len(elenchi); i++ {
			if elenchi[i].Nome == "" {
				continue
			}
			EchoString(" - " + elenchi[i].Nome)
		}
	}
	return
}
