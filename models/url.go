package models

type URL struct {
	Unique
	Name string `json:"name"`
	URL  string `json:"url"`
}
