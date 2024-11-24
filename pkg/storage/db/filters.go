package db

type SongFilter struct {
	ReleaseDate []string `json:"release_date"`
	Song        string   `json:"song"`
	Group       string   `json:"group"`
}
