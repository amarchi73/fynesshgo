package utils

import (
	"encoding/json"
	"fmt"
	"fyneapp/types"
	"io"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

var elencoFonti []string

func EseguiChiamata(d types.DatiJson) {
	if d.Pass != "" {
		fmt.Println("HOST=\"" + d.User + "@" + d.Host + "\"; export SSHPASS=$(pass sitilavoro/" + d.Pass + "); PREFIX=\"si\"; ")
		fmt.Println("echo $HOST; sshpass -e ssh -o ServerAliveInterval=5 -o ServerAliveCountMax=1 $HOST")
	} else {
		fmt.Println("HOST=\"" + d.User + "@" + d.Host + "\"; ")
		fmt.Println("echo $HOST; ssh -o ServerAliveInterval=5 -o ServerAliveCountMax=1 $HOST")
	}
}
func stripRegex(in string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(in, "")
}
func GetPass(p string) []string {
	cmd := exec.Command("pass", "sitilavoro")
	stdout, err := cmd.Output()

	if err != nil {
		//fmt.Println(err.Error())
		return nil
	}

	// Print the output
	//fmt.Println(string(stdout))
	elenco := strings.Split(string(stdout), "\n")
	for i := 0; i < len(elenco); i++ {
		elenco[i] = stripRegex(elenco[i])
	}
	sort.Strings(elenco)
	//fmt.Println(elenco)
	return elenco
}
func SalvaDati(d []types.Dati) error {

	for i := 0; i < len(d); i++ {
		if d[i].Nome == nil {
			continue
		}
		s, _ := d[i].User.Get()
		types.Data[i].User = s
		s, _ = d[i].Pass.Get()
		types.Data[i].Pass = s
		s, _ = d[i].Host.Get()
		types.Data[i].Host = s
		s, _ = d[i].Nome.Get()
		types.Data[i].Nome = s
	}
	json, _ := json.Marshal(types.Data)
	//fmt.Println(string(json))
	return os.WriteFile(types.FileJSON, json, 0666)
}
func LeggiElencoDati(fileName string) []types.DatiJson {
	f, _ := os.Open(fileName)
	jsonString, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var data []types.DatiJson
	json.Unmarshal([]byte(jsonString), &data)

	return data
}
func OrdinaElencoDati() {
	sort.Slice(types.Data, func(i, j int) bool {
		return strings.ToLower(types.Data[i].Nome) < strings.ToLower(types.Data[j].Nome)
	})
}
