package parser

import (
	"log"
	"newsbot/configs"
	"newsbot/internal/database"
	"newsbot/internal/parser/constants"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Article struct {
	Url, Image, Title, Content, DateTime string
}

// func GetArticles() {
// countPage(c, func(pagination string) {
// 	fmt.Println(pagination)
// })
// FirstPageArticles := getFirstPageArticles()

// fmt.Println(FirstPageArticles[1])
// }

func FirstPageArticles() []Article {
	c := colly.NewCollector()
	var articleList []Article
	c.OnHTML(constants.ArticleList, func(e *colly.HTMLElement) {
		e.ForEach("article", func(_ int, article *colly.HTMLElement) {
			id := article.Attr("id")
			url := constants.ArticleUrl + id
			dateTime := article.ChildAttr("time", "datetime")
			title := article.ChildText("h2 a")
			imgUrl := article.ChildAttr(constants.PathToImg, "src")
			content := article.ChildText(constants.PathToContentV2)
			if imgUrl == "" {
				imgUrl = article.ChildAttr(constants.AltPathToImg, "src")
			}
			articleList = append(articleList, Article{Title: title, Image: imgUrl, Content: content, Url: url, DateTime: dateTemplate(dateTime)})
		})
	})
	c.Visit(constants.UrlForVisit)
	return articleList
}

func AllArticles() []Article {
	var articles []Article
	count := countPage()
	for i := 1; i <= count; i++ {
		c := colly.NewCollector()
		c.OnHTML(constants.ArticleList, func(e *colly.HTMLElement) {
			e.ForEach("article", func(_ int, article *colly.HTMLElement) {
				id := article.Attr("id")
				url := constants.ArticleUrl + id
				dateTime := article.ChildAttr("time", "datetime")
				title := article.ChildText("h2 a")
				imgUrl := article.ChildAttr(constants.PathToImg, "src")
				if imgUrl == "" {
					imgUrl = article.ChildAttr(constants.AltPathToImg, "src")
				}
				content := article.ChildText(constants.PathToContentV2)
				if content == "" {
					content = article.ChildText(constants.PathToContent)
				}

				articles = append(articles, Article{Title: title, Image: imgUrl, Content: content, Url: url, DateTime: dateTemplate(dateTime)})
			})
		})
		index := strconv.Itoa(i)
		url := "page" + index + "/"
		c.Visit(constants.UrlForVisit + url)
		log.Println("page parse complete " + index)
	}
	return articles
}

func TodayArticles() []Article {
	var newArticles []Article
	articles := FirstPageArticles()
	date := currentDate()

	for _, article := range articles {
		if article.DateTime == date {
			newArticles = append(newArticles, article)
		}
	}
	log.Println("current Date: " + date)
	return newArticles
}

func currentDate() string {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	return currentDate
}

func dateTemplate(date string) string {
	strSlice := []rune(date)
	return string(strSlice[:10])
}

func InsertData() {
	cfg := configs.LoadConfig()
	db, err := database.NewDatabase(cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
		INSERT INTO articles_habr (date, title, content, url)
		VALUES ($1, $2, $3, $4)
	`
	articles := AllArticles()
	for _, article := range articles {
		_, err = db.Conn.Exec(query, article.DateTime, article.Title, article.Content, article.Url)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("data added to table!")

}

func DeleteAllRows() {
	cfg := configs.LoadConfig()
	db, err := database.NewDatabase(cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
		DELETE FROM articles_habr 
	`
	_, err = db.Conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("data deleted from table!")

}

func countPage() int {
	c := colly.NewCollector()

	var count string
	c.OnHTML("main", func(main *colly.HTMLElement) {
		count = main.DOM.Find(".tm-pagination__page-group > a").Last().Text()
	})

	c.Visit(constants.UrlForVisit)

	str := strings.TrimSpace(count)
	countPages, err := strconv.Atoi(str)
	if err != nil {
		log.Println(err)
	}
	return countPages
}
