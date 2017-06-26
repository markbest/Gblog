package api

type Category struct {
	Id           int64
	Title        string
	Articles     []SubArticle
	Sub_category []Category
}

type SubArticle struct {
	Id    int64
	Title string
}
