package ui

import (
	"groupie-tracker/api"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartApp() {
	a := app.New()
	w := a.NewWindow("Groupie Tracker")
	w.SetContent(HomePage(w))
	w.Resize(fyne.NewSize(600, 800))
	w.ShowAndRun()
	artists, err := api.FetchArtists()
	if err != nil {
		log.Fatal(err)
	}

	list := widget.NewList(
		func() int { return len(artists) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(artists[i].Name)
		},
	)

	// Quand on clique sur un artiste, afficher la fiche
	list.OnSelected = func(id widget.ListItemID) {
		artist := artists[id]
		w.SetContent(NewArtistView(artist, w))
	}

	content := container.NewVScroll(list)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 600))
	w.ShowAndRun()
}
