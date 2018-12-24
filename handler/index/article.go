package index

import (
	"blog/core"
	"blog/models"
)

type ListRequest struct {
	Page      int    `json:"page"`
	Keyword   string `json:"keyword"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// 文章列表
var ArticleList core.HandlerFunc = func(c *core.Context) core.Response {
	article := &models.Articles{}
	request := &ListRequest{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		return c.Fail(202, err)
	}
	articleResp, err := models.GetArticleList(article, request.Page, 10, request.Keyword, request.StartTime, request.EndTime)
	if err != nil {
		return c.Fail(203, err)
	}
	return c.Success(articleResp)
}
