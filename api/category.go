package api

import "time"

type Category struct {
	Id    		int64
	Title  		string
	Parent_id  	int64
	Sort       	int64
	Created_at      time.Time
	Updated_at      time.Time
}
