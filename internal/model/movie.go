package model

// Movie ...
type Movie struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Year       string `json:"year"`
	Image      string `json:"image"`
	RuntimeStr string `json:"runtimeStr"`
	Plot       string `json:"plot"`
	Genres     string `json:"genres"`
}
