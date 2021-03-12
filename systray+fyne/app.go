package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

var mainApp fyne.App
var gui fyne.Window
var showGuiCh = make(chan int)
var hideGuiCh = make(chan int)

func main() {
	initApp()
	go func() { systray.Run(runTrayMenu, nil) }()
	mainApp.Run()
}
func initApp() {
	mainApp = app.New()
	gui = mainApp.NewWindow("Hello")
	gui.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			mainApp.Quit()
		}),
	))
}
func runTrayMenu() {
	systray.SetIcon(icon.Data)
	mShow := systray.AddMenuItem("Show GUI", "Show")
	mHide := systray.AddMenuItem("Hide GUI", "Hide")
	mQuit := systray.AddMenuItem("Quit", "Quit")
	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				gui.Show()
			case <-mHide.ClickedCh:
				gui.Hide()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}
