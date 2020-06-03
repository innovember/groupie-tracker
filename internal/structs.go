package internal

type Artists struct {
	Artists []Artist
}
type Artist struct {
	ID            float32  `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	CreationDate  float32  `json:"creationDate"`
	FirstAlbum    string   `json:"firstAlbum"`
	Locations     string   `json:"locations"`
	ConcertDates  string   `json:"concertDates"`
	LocationDates string   `json:"relations"`
}

// Locations structure
type Locations struct {
	Index []SubLocations `json:"index"`
}

type SubLocations struct {
	ID        float32  `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Dates structure
type Dates struct {
	Index []SubDates `json:"index"`
}

type SubDates struct {
	ID    float32 `json:"id"`
	Dates string  `json:"dates"`
}

// Relation Structure
type Relation struct {
	Index []SubRelation `json:"index"`
}

type SubRelation struct {
	ID             float32             `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type AllRelations struct {
	AllInfo string
}
