package models

import (
	"encoding/json"
	"time"
)

type Mention struct {
	URL         string    `json:"url"`
	Text        string    `json:"text"`
	Author      Author    `json:"author"`
	Language    string    `json:"language"`
	PublishedAt time.Time `json:"published_at"`
	Source      string    `json:"source"`
}

func NewMentionFromResource(r Resource) Mention {
	return Mention{
		URL:         r.URL,
		Text:        r.Text,
		Author:      r.Author,
		Language:    r.Language,
		PublishedAt: r.PublishedAt,
		Source:      r.Source,
	}
}

func NewMentionFromJSON(j string) Mention {
	m := Mention{}
	json.Unmarshal([]byte(j), &m)
	return m
}

func (m Mention) ToJSON() string {
	b, _ := json.Marshal(m)
	return string(b)
}
