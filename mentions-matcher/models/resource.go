package models

import "time"

type Resource struct {
	URL         string    `json:"url"`
	Text        string    `json:"text"`
	Author      Author    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Source      string    `json:"source"`
}

