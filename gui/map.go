package ui

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/api"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// Coord représente les coordonnées géographiques
type Coord struct {
	Lat float64
	Lon float64
}

// Geocode adresse avec Nominatim OpenStreetMap
func Geocode(address string) (Coord, error) {
	base := "https://nominatim.openstreetmap.org/search"
	params := url.Values{}
	params.Add("q", address)
	params.Add("format", "json")
	params.Add("limit", "1")

	reqURL := base + "?" + params.Encode()

	resp, err := http.Get(reqURL)
	if err != nil {
		return Coord{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Coord{}, err
	}

	var data []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Coord{}, err
	}

	if len(data) == 0 {
		return Coord{}, fmt.Errorf("adresse introuvable : %s", address)
	}

	lat, _ := strconv.ParseFloat(data[0].Lat, 64)
	lon, _ := strconv.ParseFloat(data[0].Lon, 64)

	return Coord{Lat: lat, Lon: lon}, nil
}

// NewMapView crée le container avec la carte et marqueurs
func NewMapView() fyne.CanvasObject {
	// Placeholder : rectangle pour la carte
	rect := canvas.NewRectangle(fyne.NewColor(220, 220, 220))
	rect.SetMinSize(fyne.NewSize(600, 400))
	content := container.NewMax(rect)

	// Exemple : récupérer les locations
	locations, err := api.FetchLocations()
	if err != nil {
		log.Println("Erreur fetch locations:", err)
		return content
	}

	for _, loc := range locations {
		for _, address := range loc.Locations {
			coord, err := Geocode(address)
			if err != nil {
				log.Println("Erreur géocodage:", err)
				continue
			}

			// Créer un marqueur rouge
			marker := canvas.NewCircle(fyne.NewColor(255, 0, 0))
			marker.SetMinSize(fyne.NewSize(10, 10))

			// Position approximative (à améliorer selon l'image de carte)
			x := float32(coord.Lon+180) * 2 // simplifié pour exemple
			y := float32(90-coord.Lat) * 2
			marker.Move(fyne.NewPos(x, y))

			content.Add(marker)
		}
	}

	return content
}
