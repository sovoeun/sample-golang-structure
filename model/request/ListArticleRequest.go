package request

type ListArticleRequest struct {
	Title   string `form:"title"`
	Desc    string `form:"desc"`
	Content string `form:"content"`
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
}
