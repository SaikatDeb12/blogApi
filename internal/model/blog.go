package model

type Blog struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Author string `json:"author"`
	CreatedAt string `json:"createdAt"`
}
