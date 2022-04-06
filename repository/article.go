package repository

import (
	"errors"
	"example/sample/global"
	"example/sample/model"
	"example/sample/model/request"
	"example/sample/response"
)

func PostCreatedArticle(param *request.CreatedArticleRequest) error {
	var dataParam model.Article

	dataParam.Title = param.Title
	dataParam.Desc = param.Desc
	dataParam.Content = param.Content
	has, err := global.Db.Table(model.Article{}).InsertOne(dataParam)
	if err != nil {
		println(err.Error())
		return err
	}
	if has == 0 {
		println("does not exist")
		return errors.New("not found")
	}
	return nil

}

func PostUpdateArticle(param *request.UpdateArticleRequest) error {
	var dataParam model.Article

	dataParam.Title = param.Title
	dataParam.Desc = param.Desc
	dataParam.Content = param.Content
	has, err := global.Db.Where("title=?", param.Title).
		Update(&dataParam)
	if err != nil {
		println(err.Error())
		return err
	}
	if has == 0 {
		println("does not exist")
		return errors.New("not found")
	}
	return nil

}

func DeleteUpdateArticle(param *request.DeleteArticleRequest) error {
	var article model.Article
	println(param.Title)
	has, err := global.Db.Where("title=?", param.Title).Delete(article)
	if err != nil {
		println(err.Error())
		return err
	}
	if has == 0 {
		println(has)
		return errors.New("not found")
	}

	return nil
}

func GetArticleList(param *request.ListArticleRequest) (*response.PageResult, error) {
	var dataList response.PageResult
	var applist []*model.Article
	var dataParamList = []*response.ArticleListParam{}

	offset := 0
	if (param.Page - 1) > 0 {
		offset = (param.Page - 1) * param.Limit
	}
	session := global.Db.NewSession()
	if param.Title != "" {
		session.Where("title like ?", "%"+param.Title+"%")
	}

	if param.Desc != "" {
		session.Where("desc like ?", "%"+param.Desc+"%")
	}

	if param.Content != "" {
		session.Where("content like ?", "%"+param.Content+"%")
	}

	err := session.Table(model.Article{}).
		Limit(param.Limit, offset).
		OrderBy("title desc").
		Find(&applist)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	for _, item := range applist {
		var dataParam response.ArticleListParam
		dataParam.Title = item.Title
		dataParam.Desc = item.Desc
		dataParam.Content = item.Content
		dataParamList = append(dataParamList, &dataParam)
	}
	dataList.Total = GetArticleListTotal(param)
	dataList.Page = param.Page
	dataList.Size = param.Limit
	dataList.List = dataParamList
	return &dataList, nil
}

func GetArticleListTotal(param *request.ListArticleRequest) int64 {
	session := global.Db.NewSession()
	if param.Title != "" {
		session.Where("title like ?", "%"+param.Title+"%")
	}

	if param.Desc != "" {
		session.Where("desc like ?", "%"+param.Desc+"%")
	}

	if param.Content != "" {
		session.Where("content like ?", "%"+param.Content+"%")
	}

	total, err := session.Table(model.Article{}).
		Desc("title").
		Count()
	if err != nil {
		return total
	}
	return total
}
