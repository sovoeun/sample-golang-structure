package service

import (
	"example/sample/global"
	"example/sample/model"
	"example/sample/model/request"
	"example/sample/repository"
	"example/sample/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnAllArticles(c *gin.Context) {
	fmt.Println("Endpoint Hit: returnAllArticles")

	var Articles = []model.Article{
		model.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		model.Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	c.JSON(http.StatusOK, response.Response{
		Code: 0,
		Msg:  "",
		Data: Articles,
	})
}

func CreateArticle(r *gin.Context) {
	param := new(request.CreatedArticleRequest)
	if err := r.ShouldBind(param); err != nil {
		r.JSON(http.StatusOK, response.Response{
			Code: 100,
			Msg:  "error param1",
		})
		return
	}

	err := repository.PostCreatedArticle(param)
	if err != nil {
		global.Logger.Err(err).Send()
		r.JSON(http.StatusOK, response.Response{
			Code: 100,
			Msg:  "param err2",
		})
		return
	}
	r.JSON(http.StatusOK, response.Response{
		Code: 0,
		Msg:  "success",
	})
	fmt.Println(param.Desc)
}

func UpdateArticle(r *gin.Context) {
	param := new(request.UpdateArticleRequest)
	if err := r.ShouldBind(param); err != nil {
		r.JSON(http.StatusOK, response.Response{
			Code: 100,
			Msg:  "error param1",
		})
		return
	}

	err := repository.PostUpdateArticle(param)
	if err != nil {
		global.Logger.Err(err).Send()
		r.JSON(http.StatusOK, response.Response{
			Code: 100,
			Msg:  "param err2",
		})
		return
	}
	r.JSON(http.StatusOK, response.Response{
		Code: 0,
		Msg:  "success",
	})
	fmt.Println(param.Desc)
}

func DeleteArticle(r *gin.Context) {
	param := new(request.DeleteArticleRequest)
	if err := r.ShouldBind(param); err != nil {
		r.JSON(http.StatusOK, response.Response{
			Code: 100,
			Msg:  "error param1",
		})
		return
	}

	err := repository.DeleteUpdateArticle(param)
	if err != nil {
		global.Logger.Err(err).Send()
		r.JSON(http.StatusOK, response.Response{
			Code: 100,
			Msg:  "param err2",
		})
		return
	}
	r.JSON(http.StatusOK, response.Response{
		Code: 0,
		Msg:  "success",
	})
	fmt.Println("param.Desc")
}

func ArticleList(c *gin.Context) {
	param := new(request.ListArticleRequest)
	if err := c.ShouldBind(param); err != nil {
		c.JSON(http.StatusOK, response.Response{
			Code: 100,
			Msg:  "error param1",
		})
		return
	}

	dataList, err := repository.GetArticleList(param)
	if err != nil {
		global.Logger.Err(err).Send()
		c.JSON(http.StatusOK, response.Response{
			Code: 100,
			Msg:  "param err2",
		})
		return
	}
	response.RespOk(c, "success", dataList)
}

func HomePage(w *gin.Context) {
	fmt.Println("Endpoint Hit: homePage")
}
