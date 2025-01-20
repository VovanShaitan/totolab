package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	appIcon = "/home/vovanshaitan/Learning/GoProjects/github.com/VovanShaitan/totolab/assets/icons/flask.png"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(400, 400))

	setAppIcon(w)

	fileOpen, fileSave := createFileMenuButtons()

	menuFile := fyne.NewMenu("File", fileOpen, fileSave)

	helpAbout := fyne.NewMenuItem("About program", func() {
		fmt.Println("Program info")
	})

	menuHelp := fyne.NewMenu("Help", helpAbout)

	menuMain := fyne.NewMainMenu(menuFile, menuHelp)

	w.SetMainMenu(menuMain)

	hello := widget.NewLabel("Hello Fyne!")
	content := container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	)
	w.SetContent(content)

	w.ShowAndRun()
}

func createFileMenuButtons() (*fyne.MenuItem, *fyne.MenuItem) {
	fileOpen := fyne.NewMenuItem("Open File", func() {
		fmt.Println("File opened")
	})
	fileSave := fyne.NewMenuItem("Save File", func() {
		fmt.Println("File saved")
	})
	return fileOpen, fileSave
}

func setAppIcon(w fyne.Window) {
	ic, _ := fyne.LoadResourceFromPath(appIcon)
	w.SetIcon(ic)
}
