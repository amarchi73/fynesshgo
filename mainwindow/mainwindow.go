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
	"fyne.io/fyne/v2/widget"
	"fyneapp/types"
	"fyneapp/utils"
	"image/color"
	"log"
	"os"
)

var myApp fyne.App
var mainWindow fyne.Window
var selectElenco *widget.Select
var dd []types.Dati

func MainWindow() fyne.Window {
	myApp = app.New()
	myWindow := myApp.NewWindow("List Widget")
	mainWindow = myWindow
	dd = make([]types.Dati, len(types.Data))
	dd = append(dd, types.Dati{nil, nil, nil, nil, nil})

	list := widget.NewList(
		func() int {
			//fmt.Println("Lunghezza")
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
			/* btn := widget.NewButton("xx", func() {

			})
			box := widget.NewEntry() */
			c := container.New(layout.NewGridLayout(6), lb1, lb2, lb3, lb4, lb5, lb6)
			return c
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			dd = make([]types.Dati, len(types.Data))
			dd = append(dd, types.Dati{nil, nil, nil, nil, nil})
			//fmt.Println("Aggiorna")
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
			s.Set(types.Data[i].Nome)
			dd[i].Nome = s
			en1 := widget.NewEntryWithData(s)
			s = binding.NewString()
			s.Set(types.Data[i].User)
			dd[i].User = s
			en2 := widget.NewEntryWithData(s)
			s = binding.NewString()
			s.Set(types.Data[i].Host)
			dd[i].Host = s
			en3 := widget.NewEntryWithData(s)
			s = binding.NewString()
			s.Set(types.Data[i].Pass)
			dd[i].Pass = s
			b := binding.NewBool()
			b.Set(types.Data[i].Xwindows)
			dd[i].Xwindows = b

			//en4 := widget.NewEntryWithData(s)

			elencopass := utils.GetPass("")
			en4 := widget.NewSelect(elencopass, func(value string) {
				//log.Println("Select set to", value)
				dd[i].Pass.Set(value)
			})
			en4.SetSelected(types.Data[i].Pass)

			/* btnSave := widget.NewButton("Salva", func() {
				v, _ := dd[i].Nome.Get()
				//fmt.Println("aaaa " + strconv.Itoa(i) + " " + v)
				s, _ := dd[i].Nome.Get()
				dd[i].Nome.Set(s + ".")
				data[i].Nome = v
			}) */
			btnRun := widget.NewButton("Esegui", func() {
				fmt.Println(i)
				v, _ := dd[i].Nome.Get()
				//fmt.Println("aaaa " + strconv.Itoa(i) + " " + v)
				s, _ := dd[i].Nome.Get()
				dd[i].Nome.Set(s + ".")
				types.Data[i].Nome = v

				utils.EseguiChiamata(types.Data[i])

				myWindow.Close()
			})
			chk := widget.NewCheck("X", func(c bool) {
				dd[i].Xwindows.Set(c)
				types.Data[i].Xwindows = c

			})
			chk.SetChecked(types.Data[i].Xwindows)
			chk.Refresh()
			/* en4.OnChanged = func(s string) {
				fmt.Println("CAMBIATO " + s)
			} */
			o.(*fyne.Container).Objects[0] = en1
			o.(*fyne.Container).Objects[1] = en2
			o.(*fyne.Container).Objects[2] = en3
			o.(*fyne.Container).Objects[3] = en4
			o.(*fyne.Container).Objects[4] = chk
			o.(*fyne.Container).Objects[5] = btnRun
			//o.(*fyne.Container).Objects[5] = btnSave
		})

	os.Setenv("FOO", "1")
	//dir, _ := os.UserHomeDir()
	//fmt.Println(dir)

	//myWindow.SetContent(list)
	topDefaultBtn := widget.NewButton("Predefinito", func() {
		types.DefaultFileJson = types.FileJSON
		utils.SalvaConf("default", types.FileJSON)
	})
	top := canvas.NewText("Seleziona un elenco", color.White)
	topBtn := widget.NewButton("Seleziona", func() {
		//elencopass := []string{"uno", "due", "tre"}

		e := widget.NewEntry()
		e.SetPlaceHolder("Nuovo elenco")

		var elencopass []string
		elencopass = append(elencopass, "Nuovo")
		for i := 0; i < len(types.ElencoData); i++ {
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
			top.Refresh()
		})
		selectElenco.SetSelected(elencopass[0])
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
				}
			} else {
				ok = "NO"
			}
			dialog.ShowInformation("Creo un nuovo elenco", "Nuovo "+ok, myWindow)
		}, myWindow)
	})

	left := canvas.NewText("", color.White)
	//middle := canvas.NewText("content", color.White)

	save := widget.NewButton("Salva", func() {
		errsave := utils.SalvaDati(dd)
		if errsave == nil {
			dialog.ShowInformation("Salvataggio", "Salvato", myWindow)
		} else {
			dialog.ShowInformation("Salvataggio", "Errore "+errsave.Error(), myWindow)
		}
	})
	nuovo := widget.NewButton("Nuovo", func() {
		types.Data = append(types.Data, types.DatiJson{"", "", "", "", false})
		dd = append(dd, types.Dati{nil, nil, nil, nil, nil})
		list.Refresh()
	})
	passRefresh := widget.NewButton("Refresh", func() {
		list.Refresh()
	})
	bottom := container.New(layout.NewHBoxLayout(), save, layout.NewSpacer(), nuovo, passRefresh)
	topContainer := container.New(layout.NewHBoxLayout(), top, topBtn, topDefaultBtn)

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

	return myWindow
}

func aggiornaDati() {
	fmt.Println("Leggo " + types.FileJSON)
	types.Data = utils.LeggiElencoDati(types.FileJSON)
	utils.OrdinaElencoDati()
	selectElenco.Refresh()

	if types.DefaultFileJson == types.FileJSON {

	}
}
