package ui

import (
	"groupie-tracker/models"
	"strings"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Filters struct {
	CreationMin int
	CreationMax int
	MembersMin  int
	MembersMax  int
	Locations   []string
}

// ApplyFilters retourne la liste des artistes filtrés
func ApplyFilters(artists []models.Artist, filters Filters) []models.Artist {
	filtered := []models.Artist{}

	for _, artist := range artists {
		// Filtre date de création
		if artist.CreationDate < filters.CreationMin || artist.CreationDate > filters.CreationMax {
			continue
		}
		// Filtre nombre de membres
		if len(artist.Members) < filters.MembersMin || len(artist.Members) > filters.MembersMax {
			continue
		}
		// Filtre lieux (optionnel)
		if len(filters.Locations) > 0 {
			matched := false
			for _, loc := range filters.Locations {
				for _, aLoc := range artist.Members { // à remplacer par les lieux de concerts réels
					if strings.Contains(strings.ToLower(aLoc), strings.ToLower(loc)) {
						matched = true
					}
				}
			}
			if !matched {
				continue
			}
		}

		filtered = append(filtered, artist)
	}

	return filtered
}

// NewFiltersUI crée les widgets pour les filtres (exemple simple)
func NewFiltersUI() *widget.Box {
	creationMin := widget.NewEntry()
	creationMin.SetPlaceHolder("Année min")
	creationMax := widget.NewEntry()
	creationMax.SetPlaceHolder("Année max")

	membersMin := widget.NewEntry()
	membersMin.SetPlaceHolder("Membres min")
	membersMax := widget.NewEntry()
	membersMax.SetPlaceHolder("Membres max")

	box := container.NewVBox(
		widget.NewLabel("Filtres :"),
		creationMin, creationMax,
		membersMin, membersMax,
	)

	return box
}
