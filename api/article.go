package api

import "time"

type Article struct {
	Id         int64      `json:"id,omitempty"`
	Title      string     `json:"title,omitempty"`
	Slug       string     `json:"slug,omitempty"`
	Summary    string     `json:"summary,omitempty"`
	Views      int64      `json:"views,omitempty"`
	User       string     `json:"user,omitempty"`
	Body       string     `json:"body,omitempty"`
	Created_at *time.Time `json:"created_at,omitempty"`
	Updated_at *time.Time `json:"updated_at,omitempty"`
}
