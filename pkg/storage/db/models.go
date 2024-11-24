package db

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Song struct {
	Id          int    `json:"id"`
	Link        string `json:"link"`
	Group       `json:"group"`
	Words       string `json:"words"`
	ReleaseDate string `json:"release_date"`
	Song        string
}
