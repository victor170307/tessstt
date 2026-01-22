package api

import (
	"encoding/json"
	"net/http"

	"groupie-tracker/models" // Assure-toi que ton go.mod s'appelle bien "groupie-tracker"
)

const relationsURL = "https://groupietrackers.herokuapp.com/api/relation"

func FetchRelations() ([]models.Relation, error) {
	resp, err := http.Get(relationsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// C'EST ICI QUE ÇA CHANGE :
	// L'API renvoie un objet {"index": [...]}, on doit donc créer une structure
	// temporaire pour capturer ce champ "index".
	var apiResponse struct {
		Index []models.Relation `json:"index"`
	}

	// On décode dans cette structure temporaire
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	// On ne renvoie que la liste qui nous intéresse (le contenu de "index")
	return apiResponse.Index, nil
}
