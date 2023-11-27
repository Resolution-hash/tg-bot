package parser

import (
	_ "fmt"
	cnst "newsbot/internal/parser/constants"

	"github.com/gocolly/colly"
)

type Article struct {
	Url, Image, Name, Content, DateTime string
}

// func GetArticles() {
	// countPage(c, func(pagination string) {
	// 	fmt.Println(pagination)
	// })
	// FirstPageArticles := getFirstPageArticles()

	// fmt.Println(FirstPageArticles[1])
// }

func GetArticles() []Article {
	c := colly.NewCollector()
	var articleList []Article
	c.OnHTML(cnst.ArticleList, func(e *colly.HTMLElement) {
		e.ForEach("article", func(_ int, article *colly.HTMLElement) {
			id := article.Attr("id")
			url := cnst.ArticleUrl + id
			dateTime := article.ChildAttr("time", "datetime")
			title := article.ChildText("h2 a")
			imgUrl := article.ChildAttr(cnst.PathToImg, "src")
			content := article.ChildText(cnst.PathToText)
			if imgUrl == "" {
				imgUrl = article.ChildAttr(cnst.AltPathToImg, "src")
			}
			articleList = append(articleList, Article{Name: title, Image: imgUrl, Content: content, Url: url, DateTime: dateTime})
		})
	})
	c.Visit(cnst.UrlForVisit)
	return articleList
}

// func countPage(c *colly.Collector, callback func(string)) {
// 	c.OnHTML("main", func(main *colly.HTMLElement) {
// 		paginationList, exist := main.DOM.Find(".tm-pagination__page-group > a").Last().Attr("href")
// 		if !exist {
// 			fmt.Println("failed search pagination")
// 			return
// 		}
// 		callback(paginationList)
// 	})
// }
