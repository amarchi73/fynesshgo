package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

//var data = []string{"", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list"}

func tap() {
	log.Println("tapped")
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
func salvaDati(d []dati) error {

	for i := 0; i < len(d); i++ {
		if d[i].Nome == nil {
			continue
		}
		s, _ := d[i].User.Get()
		data[i].User = s
		s, _ = d[i].Pass.Get()
		data[i].Pass = s
		s, _ = d[i].Host.Get()
		data[i].Host = s
		s, _ = d[i].Nome.Get()
		data[i].Nome = s
	}
	json, _ := json.Marshal(data)
	//fmt.Println(string(json))
	return os.WriteFile(fileJSON, json, 0666)
}
func eseguiChiamata(d datiJson) {
	if d.Pass != "" {
		fmt.Println("HOST=\"" + d.User + "@" + d.Host + "\"; export SSHPASS=$(pass sitilavoro/" + d.Pass + "); PREFIX=\"si\"; ")
		fmt.Println("echo $HOST; sshpass -e ssh -o ServerAliveInterval=5 -o ServerAliveCountMax=1 $HOST")
	} else {
		fmt.Println("HOST=\"" + d.User + "@" + d.Host + "\"; ")
		fmt.Println("echo $HOST; ssh -o ServerAliveInterval=5 -o ServerAliveCountMax=1 $HOST")
	}
}

/*
var str [40]binding.String
var ent [40]binding.String
*/
type dati struct {
	Nome binding.String
	User binding.String
	Host binding.String
	Pass binding.String
}

//{"nome":"adler","user":"adler","host":"nomehost","pass":"password"}
type datiJson struct {
	Nome string `json:"nome"`
	User string `json:"user"`
	Host string `json:"host"`
	Pass string `json:"pass"`
}

//var dd []dati
var data []datiJson

const fileJSON = "data.json"

func main() {

	f, _ := os.Open(fileJSON)
	jsonString, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//s := string(jsonString)
	//fmt.Println(s)
	json.Unmarshal([]byte(jsonString), &data)
	sort.Slice(data, func(i, j int) bool {
		return data[i].Nome < data[j].Nome
	})
	if len(os.Args) > 1 {
		nome := os.Args[1]
		if nome == "-t" {
			fmt.Println("echo \"ssh -o ServerAliveInterval=5 -o ServerAliveCountMax=1 <HOST>\"")
			return
		}
		//fmt.Println(nome)
		for i := 0; i < len(data); i++ {
			if nome == data[i].Nome {
				//fmt.Println("TROVATO " + nome)
				eseguiChiamata(data[i])
			}
		}
		return
	}

	//fmt.Println(data)
	dd := make([]dati, len(data))
	dd = append(dd, dati{nil, nil, nil, nil})
	//fmt.Println(dd)

	myApp := app.New()
	myWindow := myApp.NewWindow("List Widget")
	list := widget.NewList(
		func() int {
			return len(data) + 1
		},
		func() fyne.CanvasObject {
			//return widget.NewLabel("template")
			lb1 := widget.NewLabel("Nome")
			lb2 := widget.NewLabel("User")
			lb3 := widget.NewLabel("Host")
			lb4 := widget.NewLabel("Pass")
			//lb5 := widget.NewLabel("Salva")
			lb6 := widget.NewLabel("Esegui")
			/* btn := widget.NewButton("xx", func() {

			})
			box := widget.NewEntry() */
			c := container.New(layout.NewGridLayout(5), lb1, lb2, lb3, lb4, lb6)
			return c
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			/* lb := widget.NewLabel("template")
			btn := widget.NewButton("xx", func() {
				fmt.Println("aaaa")
			})
			container.Ob
			o.(*container).New(layout.NewGridLayout(2))
			o.(*widget.Label).SetText(data[i]) */
			if i == 0 {
				return
			}
			i--

			s := binding.NewString()
			s.Set(data[i].Nome)
			dd[i].Nome = s
			en1 := widget.NewEntryWithData(s)
			s = binding.NewString()
			s.Set(data[i].User)
			dd[i].User = s
			en2 := widget.NewEntryWithData(s)
			s = binding.NewString()
			s.Set(data[i].Host)
			dd[i].Host = s
			en3 := widget.NewEntryWithData(s)
			s = binding.NewString()
			s.Set(data[i].Pass)
			dd[i].Pass = s
			//en4 := widget.NewEntryWithData(s)

			elencopass := GetPass("")
			en4 := widget.NewSelect(elencopass, func(value string) {
				//log.Println("Select set to", value)
				dd[i].Pass.Set(value)
			})
			en4.SetSelected(data[i].Pass)

			/* btnSave := widget.NewButton("Salva", func() {
				v, _ := dd[i].Nome.Get()
				//fmt.Println("aaaa " + strconv.Itoa(i) + " " + v)
				s, _ := dd[i].Nome.Get()
				dd[i].Nome.Set(s + ".")
				data[i].Nome = v
			}) */
			btnRun := widget.NewButton("Esegui", func() {
				v, _ := dd[i].Nome.Get()
				//fmt.Println("aaaa " + strconv.Itoa(i) + " " + v)
				s, _ := dd[i].Nome.Get()
				dd[i].Nome.Set(s + ".")
				data[i].Nome = v

				eseguiChiamata(data[i])

				myWindow.Close()
			})

			/* en4.OnChanged = func(s string) {
				fmt.Println("CAMBIATO " + s)
			} */
			o.(*fyne.Container).Objects[0] = en1
			o.(*fyne.Container).Objects[1] = en2
			o.(*fyne.Container).Objects[2] = en3
			o.(*fyne.Container).Objects[3] = en4
			o.(*fyne.Container).Objects[4] = btnRun
			//o.(*fyne.Container).Objects[5] = btnSave
		})

	os.Setenv("FOO", "1")
	//dir, _ := os.UserHomeDir()
	//fmt.Println(dir)

	//myWindow.SetContent(list)

	top := canvas.NewText("Seleziona un elenco", color.White)
	topBtn := widget.NewButton("Seleziona", func() {
		elencopass := []string{"uno", "due", "tre"}
		selectElenco := widget.NewSelect(elencopass, func(value string) {
			//log.Println("Select set to", value)
		})
		selectElenco.SetSelected(elencopass[1])
		e := widget.NewEntry()
		e.SetPlaceHolder("Nuovo elenco")
		c := container.New(layout.NewVBoxLayout(), selectElenco, e)
		dialog.ShowCustomConfirm("Seleziona un elenco", "nuovo", "chiudi", c, func(ret bool) {
			ok := ""
			if ret {
				ok = "SI"
			} else {
				ok = "NO"
			}
			dialog.ShowInformation("Creo un nuovo elenco", "Nuovo "+ok, myWindow)
		}, myWindow)
	})
	left := canvas.NewText("", color.White)
	//middle := canvas.NewText("content", color.White)

	save := widget.NewButton("Salva", func() {
		errsave := salvaDati(dd)
		if errsave == nil {
			dialog.ShowInformation("Salvataggio", "Salvato", myWindow)
		} else {
			dialog.ShowInformation("Salvataggio", "Errore "+errsave.Error(), myWindow)
		}
	})
	nuovo := widget.NewButton("Nuovo", func() {
		data = append(data, datiJson{"", "", "", ""})
		dd = append(dd, dati{nil, nil, nil, nil})
		list.Refresh()
	})
	passRefresh := widget.NewButton("Refresh", func() {
		list.Refresh()
	})
	bottom := container.New(layout.NewHBoxLayout(), save, layout.NewSpacer(), nuovo, passRefresh)
	topContainer := container.New(layout.NewHBoxLayout(), top, topBtn)

	content := container.New(layout.NewBorderLayout(topContainer, bottom, left, nil),
		topContainer, left, list, bottom)
	myWindow.SetContent(content)

	ico, _ := fyne.LoadResourceFromPath("/home/adler/prove/fyne/icona.png")
	myWindow.SetIcon(ico)
	size := fyne.Size{1000, 500}
	myWindow.Resize(size)
	myWindow.CenterOnScreen()

	/* dialog.ShowFileOpen(func(u fyne.URIReadCloser, e error) {
		fmt.Println(u.URI().String())
	}, myWindow) */

	myWindow.ShowAndRun()

}
