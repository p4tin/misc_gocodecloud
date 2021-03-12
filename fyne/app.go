package main

import (
	"encoding/json"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	jsonc "github.com/nwidger/jsoncolor"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Two Way")

	str := binding.NewString()
	str.Set("<paste json here and press format>")

	entry := widget.NewEntryWithData(str)
	entry.MultiLine = true

	lab := widget.NewLabel("")

	content := widget.NewButton("Format", func() {
		var inter map[string]interface{}
		jsonStr, err := str.Get()
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal([]byte(jsonStr), &inter)
		if err != nil {
			panic(err)
		}

		byt, err := jsonc.MarshalIndent(inter, "", "    ")
		if err != nil {
			panic(err)
		}
		lab.SetText(string(byt))
	})

	w.SetContent(container.NewVBox(
		entry,
		content,
		lab,
	))

	w.ShowAndRun()
}
