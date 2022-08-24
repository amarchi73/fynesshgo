package mainwindow

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyneapp/types"
	"fyneapp/utils"
	"image/color"
	"log"
	"os"
	"strings"
)

type wLista struct {
	en1    *widget.Entry
	en2    *widget.Entry
	en3    *widget.Entry
	sel    *widget.Select
	chk    *widget.Check
	btn    *widget.Button
	btndel *widget.Button
}

var myApp fyne.App
var mainWindow fyne.Window
var selectElenco *widget.Select
var datiLista []types.Dati
var elemLista []wLista

var list *widget.List

func MainWindow() fyne.Window {
	myApp = app.New()
	myWindow := myApp.NewWindow("List Widget")
	mainWindow = myWindow

	ico, _ := fyne.LoadResourceFromPath("/home/adler/prove/fyne/icona.png")
	myWindow.SetIcon(ico)
	size := fyne.Size{1000, 500}
	myWindow.Resize(size)
	myWindow.CenterOnScreen()

	caricaDatiForm()

	list = widget.NewList(
		func() int {
			//fmt.Println(time.Now(), "Lunghezza", len(types.Data)+1)
			return len(types.Data) + 1
		},
		func() fyne.CanvasObject {
			//fmt.Println("Creazione")

			//return widget.NewLabel("template")
			lb1 := widget.NewLabel("Nome")
			lb2 := widget.NewLabel("User")
			lb3 := widget.NewLabel("Host")
			lb4 := widget.NewLabel("Pass")
			lb5 := widget.NewLabel("X")
			//lb5 := widget.NewLabel("Salva")
			lb6 := widget.NewLabel("Esegui")
			lb7 := widget.NewLabel("Elimina")
			/* btn := widget.NewButton("xx", func() {

			})
			box := widget.NewEntry() */
			gridLayout := layout.NewGridLayout(7)
			sizeGrid := gridLayout.MinSize([]fyne.CanvasObject{lb1, lb2, lb3, lb4, lb5, lb6, lb7})
			gridLayout.Layout([]fyne.CanvasObject{lb1, lb2, lb3, lb4, lb5, lb6, lb7}, sizeGrid)
			c := container.New(gridLayout, lb1, lb2, lb3, lb4, lb5, lb6, lb7)

			return c
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {

			//fmt.Println(time.Now(), "Aggiorna")
			if i == 0 {
				return
			}
			i--

			o.(*fyne.Container).Objects[0] = elemLista[i].en1
			o.(*fyne.Container).Objects[1] = elemLista[i].en2
			o.(*fyne.Container).Objects[2] = elemLista[i].en3
			o.(*fyne.Container).Objects[3] = elemLista[i].sel
			o.(*fyne.Container).Objects[4] = elemLista[i].chk
			o.(*fyne.Container).Objects[5] = elemLista[i].btndel
			o.(*fyne.Container).Objects[6] = elemLista[i].btn

		},
	)

	os.Setenv("FOO", "1")

	topDefaultBtn := widget.NewButton("Predefinito", func() {
		types.DefaultFileJson = types.FileJSON
		utils.SalvaConf("default", types.FileJSON)
	})

	curPass := strings.ReplaceAll(types.FileJSON, ".json", "")
	fmt.Println("curpass", curPass)
	topTxt := "Seleziona un elenco"
	if curPass != "" {
		topTxt = "Elenco Selezionato " + curPass
	}
	top := canvas.NewText(topTxt, color.White)
	topBtn := widget.NewButton("Seleziona", func() {
		//elencopass := []string{"uno", "due", "tre"}

		e := widget.NewEntry()
		e.SetPlaceHolder("Nuovo elenco")

		var elencopass []string
		elencopass = append(elencopass, "Nuovo")
		indPass := 0
		for i := 0; i < len(types.ElencoData); i++ {
			if curPass == types.ElencoData[i].Nome {
				indPass = i
			}
			elencopass = append(elencopass, types.ElencoData[i].Nome)
		}
		selectElenco = widget.NewSelect(elencopass, func(value string) {
			log.Println("Select set to", value, selectElenco.SelectedIndex())
			if selectElenco.SelectedIndex()-1 >= 0 {
				types.FileJSON = types.ElencoData[selectElenco.SelectedIndex()-1].Path
				top.Text = "Elenco Selezionato " + value
			} else {
				types.FileJSON = ""
				top.Text = "Seleziona un elenco"
			}
			aggiornaDati()
			fmt.Println("SELEZIONATO")
			top.Refresh()
		})
		selectElenco.SetSelected(types.ElencoData[indPass].Path)
		c := container.New(layout.NewVBoxLayout(), selectElenco, e)
		dialog.ShowCustomConfirm("Seleziona un elenco", "nuovo/usa", "chiudi", c, func(ret bool) {
			ok := ""
			if ret {
				ok = ""
				for i := 0; i < len(types.ElencoData); i++ {
					if types.ElencoData[i].Nome == e.Text {
						ok = "ERRORE, nome già presente: " + e.Text
						break
					}
				}
				if ok == "" && e.Text != "" {
					types.ElencoData = append(types.ElencoData, types.ElencoDatiJson{e.Text, utils.Pulisci(e.Text) + ".json"})
					utils.SalvaElenchi()
					ok = "SI " + e.Text
					types.FileJSON = utils.Pulisci(e.Text) + ".json"
					top.Text = "Elenco Selezionato " + e.Text
					top.Refresh()
					aggiornaDati()

					dialog.ShowInformation("Creo un nuovo elenco", "Nuovo "+ok, myWindow)
				}
			} else {
				ok = "NO"
			}
		}, myWindow)
	})

	left := canvas.NewText("", color.White)
	//middle := canvas.NewText("content", color.White)

	save := widget.NewButton("Salva", func() {
		errsave := utils.SalvaDati(datiLista)
		if errsave == nil {
			dialog.ShowInformation("Salvataggio", "Salvato", myWindow)
		} else {
			dialog.ShowInformation("Salvataggio", "Errore "+errsave.Error(), myWindow)
		}
	})
	nuovo := widget.NewButton("Nuovo", func() {
		types.Data = append(types.Data, types.DatiJson{"", "", "", "", false})
		//datiLista = append(datiLista, types.Dati{nil, nil, nil, nil, nil})

		elemLista = append(elemLista, wLista{nil, nil, nil, nil, nil, nil, nil})
		caricaDatiForm()
		list.Refresh()
	})
	passRefresh := widget.NewButton("Refresh", func() {
		caricaDatiForm()
		list.Refresh()
	})
	bottom := container.New(layout.NewHBoxLayout(), save, layout.NewSpacer(), nuovo, passRefresh)
	topContainer := container.New(layout.NewHBoxLayout(), top, topBtn, topDefaultBtn)

	content := container.New(layout.NewBorderLayout(topContainer, bottom, left, nil),
		topContainer, left, list, bottom)
	myWindow.SetContent(content)

	/* dialog.ShowFileOpen(func(u fyne.URIReadCloser, e error) {
		fmt.Println(u.URI().String())
	}, myWindow) */

	return myWindow
}

func caricaDatiForm() {
	elemLista = make([]wLista, len(types.Data))
	datiLista = make([]types.Dati, len(types.Data))
	datiLista = append(datiLista, types.Dati{nil, nil, nil, nil, nil})

	for i := 0; i < len(types.Data); i++ {
		s := binding.NewString()
		s.Set(types.Data[i].Nome)
		datiLista[i].Nome = s
		//en1 := widget.NewEntryWithData(s)
		s = binding.NewString()
		s.Set(types.Data[i].User)
		datiLista[i].User = s
		//en2 := widget.NewEntryWithData(s)
		s = binding.NewString()
		s.Set(types.Data[i].Host)
		datiLista[i].Host = s
		//en3 := widget.NewEntryWithData(s)
		s = binding.NewString()
		s.Set(types.Data[i].Pass)
		datiLista[i].Pass = s
		b := binding.NewBool()
		b.Set(types.Data[i].Xwindows)
		datiLista[i].Xwindows = b

		// Per una questione di scope serve questo accrocchio
		ii := i

		elemLista[i].en1 = widget.NewEntryWithData(datiLista[i].Nome)
		elemLista[i].en2 = widget.NewEntryWithData(datiLista[i].User)
		elemLista[i].en3 = widget.NewEntryWithData(datiLista[i].Host)
		pswList := utils.GetPass(types.FileJSON)
		elemLista[i].sel = widget.NewSelect(pswList, func(value string) {
			i := ii
			log.Println("Select set to", value)
			datiLista[i].Pass.Set(value)
		})
		elemLista[i].sel.SetSelected(types.Data[i].Pass)

		elemLista[i].btn = widget.NewButton("Esegui", func() {
			i := ii
			fmt.Println("====", i, datiLista)
			v, _ := datiLista[i].Nome.Get()
			//fmt.Println("aaaa " + strconv.Itoa(i) + " " + v)
			s, _ := datiLista[i].Nome.Get()
			datiLista[i].Nome.Set(s + ".")
			types.Data[i].Nome = v
			utils.EseguiChiamata(types.Data[i])
			mainWindow.Close()
		})
		elemLista[i].btn.Icon = theme.ComputerIcon()
		elemLista[i].btndel = widget.NewButton("Del", func() {
			i := ii
			fmt.Println("====", i, datiLista)
			datiLista = append(datiLista[:i], datiLista[i+1:]...)
			types.Data = append(types.Data[:i], types.Data[i+1:]...)
			caricaDatiForm()
			list.Refresh()

			/*v, _ := datiLista[i].Nome.Get()
			s, _ := datiLista[i].Nome.Get()
			datiLista[i].Nome.Set(s + ".")
			types.Data[i].Nome = v
			utils.EseguiChiamata(types.Data[i])
			mainWindow.Close()*/
		})
		elemLista[i].btndel.Icon = theme.DeleteIcon()
		elemLista[i].chk = widget.NewCheck("X", func(c bool) {
			i := ii
			datiLista[i].Xwindows.Set(c)
			types.Data[i].Xwindows = c

		})
		elemLista[i].chk.SetChecked(types.Data[i].Xwindows)
	}
}
func aggiornaDati() {
	if types.FileJSON == "" {
		return
	}
	fmt.Println("Leggo " + types.FileJSON)
	types.Data = utils.LeggiElencoDati(types.FileJSON)
	fmt.Println("====**====")
	utils.OrdinaElencoDati()

	caricaDatiForm()
	list.Refresh()

	if types.DefaultFileJson == types.FileJSON {

	}
}
