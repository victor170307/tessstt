package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchLocations(artistID int) ([]string, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", artistID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Locations []string `json:"locations"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Locations, nil
}
