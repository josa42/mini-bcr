package models

import (
	"encoding/json"
	"time"
)

type Author struct {
	Name     string
	Username string
}

type Resource struct {
	URL         string    `json:"url"`
	Text        string    `json:"text"`
	Author      Author    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Source      string    `json:"source"`
}

func (r Resource) Encode() string {
	str, _ := json.Marshal(r)
	return string(str)
}

