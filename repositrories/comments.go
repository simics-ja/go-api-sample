package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/simics-ja/go-api-sample/models"
)

// 新規投稿をDBにinsertする関数
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`

	newComment := models.Comment{
		ArticleID: comment.ArticleID,
		Message: comment.Message,
	}

	result, err := db.Exec(sqlStr, newComment.ArticleID, newComment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

// 指定IDの記事についたコメント一覧を取得する関数
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment,0)
	for rows.Next() {
		var comment models.Comment
		var createdAt sql.NullTime
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdAt)

		if createdAt.Valid {
			comment.CreatedAt = createdAt.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}