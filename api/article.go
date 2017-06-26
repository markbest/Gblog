package api

import "time"

type Article struct {
	Id         int64
	Title      string
	Slug       string
	Summary    string
	Views      int64
	User       string
	Body       string
	Created_at time.Time
	Updated_at time.Time
}
