package mainwindow

import (
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
	"os"
)

var myApp fyne.App
var mainWindow fyne.Window

func MainWindow(dd []types.Dati) fyne.Window {
	myApp = app.New()
	myWindow := myApp.NewWindow("List Widget")
	mainWindow = myWindow

	list := widget.NewList(
		func() int {
			return len(types.Data) + 1
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
				v, _ := dd[i].Nome.Get()
				//fmt.Println("aaaa " + strconv.Itoa(i) + " " + v)
				s, _ := dd[i].Nome.Get()
				dd[i].Nome.Set(s + ".")
				types.Data[i].Nome = v

				utils.EseguiChiamata(types.Data[i])

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
		errsave := utils.SalvaDati(dd)
		if errsave == nil {
			dialog.ShowInformation("Salvataggio", "Salvato", myWindow)
		} else {
			dialog.ShowInformation("Salvataggio", "Errore "+errsave.Error(), myWindow)
		}
	})
	nuovo := widget.NewButton("Nuovo", func() {
		types.Data = append(types.Data, types.DatiJson{"", "", "", ""})
		dd = append(dd, types.Dati{nil, nil, nil, nil})
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

	return myWindow
}
