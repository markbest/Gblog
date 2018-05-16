package models

import "time"

type ApiArticle struct {
	Id        int64      `json:"id"`
	Title     string     `json:"title"`
	Slug      string     `json:"slug,omitempty"`
	Summary   string     `json:"summary,omitempty"`
	Views     int64      `json:"views,omitempty"`
	User      string     `json:"user,omitempty"`
	Body      string     `json:"body,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ApiCategory struct {
	Id          int64         `json:"id"`
	Title       string        `json:"title"`
	Articles    []ApiArticle  `json:"articles,omitempty"`
	SubCategory []ApiCategory `json:"sub_category,omitempty"`
}

type ApiSearch struct {
	Id        int64     `json:"id"`
	Url       string    `json:"url,omitempty"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
