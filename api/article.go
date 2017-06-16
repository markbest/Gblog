package api

import "time"

type Article struct {
	Id    		int64
	Title  		string
	Slug  		string
	Summary		string
	Body		string
	Image           string
	Views		int64
	User_id		int64
	Created_at      time.Time
	Updated_at      time.Time
	Category	Category
}