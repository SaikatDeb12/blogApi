package model

import "time"

type Blog struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Author string `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}
