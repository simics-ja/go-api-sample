package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/simics-ja/go-api-sample/handlers"
// )

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/simics-ja/go-api-sample/models"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// * Insert a row
	article := models.Article{
		Title: "insert test",
		Contents: "Can I insert data correctly?",
		UserName: "saki",
	}

	const sqlStr =`
		insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	// * Read a row
	// articleID := 1
	// const sqlStr = `
	// 	select *
	// 	from articles
	// 	where article_id = ?;
	// `

	// row := db.QueryRow(sqlStr, articleID)
	// if err := row.Err(); err != nil {
	// 	fmt.Println(err)
	// }

	// var article models.Article
	// var createdTime sql.NullTime

	// err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if createdTime.Valid {
	// 	article.CreatedAt = createdTime.Time
	// }

	// fmt.Printf("%+v\n", article)

	// * Original code
	// r := mux.NewRouter()

	// r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	// r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
  // r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
  // r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
  // r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	// log.Println("server start at port 8080")
	// log.Fatal(http.ListenAndServe(":8080", r))
}