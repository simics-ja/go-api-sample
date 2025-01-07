package testdata

import "github.com/simics-ja/go-api-sample/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID: 1,
		Title: "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum: 2,
	},
	models.Article{
		ID: 2,
		Title: "2nd",
		Contents: "Second blog post",
		UserName: "saki",
		NiceNum: 4,
	},
}