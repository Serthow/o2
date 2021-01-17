package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"o2/snes"
)

var (
	a     fyne.App
	w     fyne.Window
	snesC chan snes.DeviceDescriptor = make(chan snes.DeviceDescriptor)
)

func main() {
	a = app.NewWithID("o2")

	go appMain()

	w = a.NewWindow("O2")
	setContent(w)
	w.SetMaster()
	w.ShowAndRun()
}

func setContent(w fyne.Window) {
	menu := widget.NewList(
		func() int {
			return len(Screens)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Item")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(Screens[id].Label())
		},
	)

	content := container.NewMax()
	title := widget.NewLabel("")
	menu.OnSelected = func(id widget.ListItemID) {
		screen := Screens[id]
		title.SetText(screen.Label())
		v := screen.View(w)
		if v != nil {
			content.Objects = []fyne.CanvasObject{v}
		} else {
			content.Objects = []fyne.CanvasObject{}
		}
		content.Refresh()
	}

	right := container.NewBorder(
		title,
		nil,
		nil,
		nil,
		content)

	split := container.NewHSplit(
		menu,
		right)
	split.Offset = 0.2

	w.SetContent(split)

	menu.Select(0)
	w.Resize(fyne.NewSize(640, 480))
}

func appMain() {
	for {
		select {
		case dev := <-snesC:
			fmt.Println(dev.DisplayName())
			break
		}
	}

	//time.NewTicker(16700 * time.Microsecond)
}
