package postmanager

import (
	"log"
	"newsbot/configs"
	"newsbot/internal/database"
	"newsbot/internal/models"
)

type Post struct {
	Url, Content, DateTime string
}

// func FirstPagePosts() []Post {
// 	return posts(parser.FirstPageArticles())
// }

// // func TodayPosts() []Post {
// // 	return posts(parser.TodayArticles())
// // }

func RandomPost() []models.Article {
	cfg := configs.LoadConfig()
	db, err := database.NewDatabase(cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	article := database.SelectRandomArticle(db.Conn)
	log.Println("SelectRandomArticle end")

	return []models.Article{article}
}

func TodayPosts() []models.Article {
	cfg := configs.LoadConfig()
	db, err := database.NewDatabase(cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	article := database.SelectTodayPosts(db.Conn)
	log.Println("article ", article)
	return []models.Article{article}
}

// func posts(articles []parser.Article) []Post {
// 	var post []Post
// 	for _, article := range articles {
// 		post = append(post, Post{Url: article.Url, Content: article.Content, DateTime: article.DateTime})
// 	}
// 	return post
// }
