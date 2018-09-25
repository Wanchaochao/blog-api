package admin

import (
	"blog/core"
	"blog/models"
	"strconv"
	"github.com/ilibs/gosql"
)

// 文章列表
var ArticleList core.HandlerFunc = func(c *core.Context) core.Response {
	article := &models.Articles{}
	page,err := strconv.Atoi(c.DefaultQuery("page","1"))
	if err != nil {
		return c.Fail(202,err)
	}
	keyword := c.DefaultQuery("keyword","")
	startTime := c.DefaultQuery("start_time","")
	endTime := c.DefaultQuery("start_time","")
	articleList,err := models.GetArticleList(article,page,10,keyword,startTime,endTime)
	if err != nil {
		return c.Fail(203,err)
	}
	return c.Success(articleList)
}

// 创建文章
var CreateArticle core.HandlerFunc = func(c *core.Context) core.Response {

	article := &models.Articles{}

	if err := c.ShouldBindJSON(article); err != nil {
		return c.Fail(301,err)
	}
	if article.CategoryId == "" || article.Title == "" || article.Description =="" || article.Author =="" {
		return c.Fail(202,"params are not permitted")
	}
	if len(article.Content) < 50 {
		return c.Fail(204,"the content of article is too little")
	}
	cate := &models.Category{}
	err := gosql.Model(cate).Where("id = ?",article.CategoryId).Get()
	if err != nil {
		return c.Fail(203,"the article category is not existed!")
	}

	if _, err := gosql.Model(article).Create(); err != nil {
		return c.Fail(204,err)
	}
	return c.Success(nil)
}


// 删除文章
var DeleteArticle core.HandlerFunc = func(c *core.Context) core.Response {
	article := &models.Articles{}
	id := c.DefaultQuery("id","")
	if id == "" {
		return c.Fail(203,"missing param article id")
	}
	if _ ,err := gosql.Model(article).Where("id = ?", id).Delete() ; err != nil{
		return c.Fail(301,err)
	}
	return c.Success("删除成功")
}

