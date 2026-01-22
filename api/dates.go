package api

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
)

const datesURL = "https://groupietrackers.herokuapp.com/api/dates"

func FetchDates() ([]models.Date, error) {
	resp, err := http.Get(datesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dates []models.Date
	err = json.NewDecoder(resp.Body).Decode(&dates)
	return dates, err
}
