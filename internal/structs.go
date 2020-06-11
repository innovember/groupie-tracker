package internal

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
	ID float32 `json:"id"`
	// DatesLocations map[string]string `json:"DatesLocations"`
	Location string `json:"location"`
	Date     string `json:"date"`
}

type AllRelationsData struct {
	Draw            int            `json:"draw"`
	RecordsTotal    int            `json:"recordsTotal"`
	RecordsFiltered int            `json:"recordsFiltered"`
	Data            []AllRelations `json:"data"`
}

// Artist personal info
type ArtistInfo struct {
	ID            float32
	Image         string
	Name          string
	Members       []string
	CreationDate  float32
	FirstAlbum    string
	LocationDates SubRelation
}

// Search result structure
type Search struct {
	Artists   []Artist
	Relations Relation
}

// Map struct

type MapData struct {
	LocationsArr []string
	Count        int
}
