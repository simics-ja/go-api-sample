package repositories_test

import (
	"testing"

	"github.com/simics-ja/go-api-sample/models"
	"github.com/simics-ja/go-api-sample/repositories"

	_ "github.com/go-sql-driver/mysql"
)

// SelectCommentList 関数のテスト
func TestSelectCommentList(t *testing.T) {
	articleID := 1
	expectNum := 2
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectNum {
		t.Errorf("got %d but want %d\n", num, expectNum)
	}
}

// InsertComment 関数のテスト
func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message: "CommentInsertTest",
	}

	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?
		`
		testDB.Exec(sqlStr, comment.Message)
	})
}