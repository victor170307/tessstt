package models

// Structure principale pour un artiste
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`    // URL vers locations
	ConcertDates string   `json:"concertDates"` // URL vers dates
	Relations    string   `json:"relations"`    // URL vers relations
}

// Structure pour stocker les relations (Villes + Dates)
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Structure pour la réponse de l'API relation
type RelationIndex struct {
	Index []Relation `json:"index"`
}

// Structure pour la géolocalisation (Map API)
type GeoLocation struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// Structure pour la barre de recherche (Suggestions)
type SearchResult struct {
	Text     string `json:"text"`
	Type     string `json:"type"` // "artist", "member", "location", etc.
	ArtistID int    `json:"artist_id"`
}
