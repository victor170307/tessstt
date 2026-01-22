package ui

import (
	"fmt"
	"groupie-tracker/models"
	"strings"

	"fyne.io/fyne/v2/widget"
)

type SearchResult struct {
	Text string
	Type string
}

// SearchArtists filtre la liste des artistes selon le texte saisi
func SearchArtists(artists []models.Artist, query string) []SearchResult {
	query = strings.ToLower(query)
	results := []SearchResult{}

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			results = append(results, SearchResult{Text: artist.Name, Type: "Artist"})
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				results = append(results, SearchResult{Text: member, Type: "Member"})
			}
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			results = append(results, SearchResult{Text: artist.FirstAlbum, Type: "First Album"})
		}
		if strings.Contains(strings.ToLower(fmt.Sprint(artist.CreationDate)), query) {
			results = append(results, SearchResult{Text: fmt.Sprint(artist.CreationDate), Type: "Creation Year"})
		}
	}

	return results
}

// NewSearchBar crée la barre de recherche pour la GUI
func NewSearchBar(artists []models.Artist, onSelect func(models.Artist)) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Rechercher un artiste ou membre...")

	entry.OnChanged = func(text string) {
		results := SearchArtists(artists, text)
		// Ici on peut afficher les suggestions (à intégrer plus tard dans un container)
		_ = results
	}

	return entry
}
