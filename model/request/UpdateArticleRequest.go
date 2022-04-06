package request

type UpdateArticleRequest struct {
	Title   string `form:"title" binding:"required"`
	Desc    string `form:"desc"`
	Content string `form:"content"`
}
