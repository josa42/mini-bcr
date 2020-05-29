package models

import (
	"encoding/json"
	"time"
)

type Author struct {
	Name     string
	Username string
}

type Mention struct {
	URL         string    `json:"url"`
	Text        string    `json:"text"`
	Author      Author    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Source      string    `json:"source"`
}

func NewMentionFromResource(r Resource) Mention {
	return Mention{
		URL:         r.URL,
		Text:        r.Text,
		Author:      r.Author,
		PublishedAt: r.PublishedAt,
		Source:      r.Source,
	}
}

func (m *Mention) ToJSON() string {
	b, _ := json.Marshal(m)
	return string(b)
}

