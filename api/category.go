package api

type Category struct {
	Id           int64      `json:"id,omitempty"`
	Title        string     `json:"title,omitempty"`
	Articles     []Article  `json:"articles,omitempty"`
	Sub_category []Category `json:"sub_category,omitempty"`
}
