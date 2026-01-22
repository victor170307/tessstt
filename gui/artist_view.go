package ui

import (
	"fmt"
	"groupie-tracker/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// NewArtistView crée un container pour afficher les détails d'un artiste
func NewArtistView(artist models.Artist, w fyne.Window) fyne.CanvasObject {
	img := canvas.NewImageFromURI(nil)
	if artist.Image != "" {
		uri, err := fyne.LoadResourceFromURLString(artist.Image)
		if err == nil {
			img = canvas.NewImageFromResource(uri)
			img.FillMode = canvas.ImageFillContain
			img.SetMinSize(fyne.NewSize(200, 200))
		}
	}

	info := widget.NewLabel(
		"Nom: " + artist.Name +
			"\nCréation: " + fmt.Sprint(artist.CreationDate) +
			"\nPremier Album: " + artist.FirstAlbum +
			"\nMembres: " + fmt.Sprint(artist.Members),
	)
	info.Wrapping = fyne.TextWrapWord

	back := widget.NewButton("Retour", func() {
		w.Content().Refresh()
	})

	content := container.NewVBox(img, info, back)
	return content
}
