package main

import (
	"fmt"
	"fyneapp/mainwindow"
	"fyneapp/types"
	_ "fyneapp/types"
	"fyneapp/utils"
	"log"
	"os"
	"strings"
)

//var data = []string{"", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list"}

func tap() {
	log.Println("tapped")
}

/*
var str [40]binding.String
var ent [40]binding.String
*/

//var dd []dati

func main() {

	/*utils.LeggiElencoPass("fava")
	fmt.Println("=======")
	utils.LeggiElencoPass("prova")
	fmt.Println("=======")*/
	ok := utils.ConfigurazionePresente()
	if ok == types.ConfigNonValido {
		return
	}
	/*f, _ := os.Open(types.FileJSON)
	jsonString, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//s := string(jsonString)
	//fmt.Println(s)
	json.Unmarshal([]byte(jsonString), &types.Data)*/

	utils.CaricaElenchi()

	/*sort.Slice(types.Data, func(i, j int) bool {
		return types.Data[i].Nome < types.Data[j].Nome
	})*/
	if len(os.Args) > 1 {
		nome := os.Args[1]
		if nome == "-h" {
			utils.StampaSintassi()
			return
		}
		if nome == "-t" {
			fmt.Println("echo \"ssh -o ServerAliveInterval=5 -o ServerAliveCountMax=1 <HOST>\"")
			return
		}
		if strings.Contains(nome, "-l") {
			utils.ListPass(nome)
			return
		}
		if strings.Contains(nome, "-e=") && len(os.Args) > 2 {
			elenco := strings.ReplaceAll(nome, "-e=", "")
			types.Data = utils.LeggiElencoDati(elenco)
			utils.OrdinaElencoDati()
			nome = os.Args[2]
		} else {
			types.Data = utils.LeggiElencoDati(types.FileJSON)
			utils.OrdinaElencoDati()
		}
		//fmt.Println(nome)
		for i := 0; i < len(types.Data); i++ {
			if nome == types.Data[i].Nome {
				//fmt.Println("TROVATO " + nome)
				utils.EseguiChiamata(types.Data[i])
				return
			}
		}
		fmt.Println("echo \"Non trovato\"")
		return
	}

	types.Data = utils.LeggiElencoDati(types.FileJSON)
	utils.OrdinaElencoDati()
	//fmt.Println(data)
	/* dd := make([]types.Dati, len(types.Data))
	dd = append(dd, types.Dati{nil, nil, nil, nil}) */
	//fmt.Println(dd)

	myWindow := mainwindow.MainWindow()

	myWindow.ShowAndRun()

}
