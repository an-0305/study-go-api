package repositories_test

import (
	"testing"

	"github.com/an-0305/study-go-api/models"
	"github.com/an-0305/study-go-api/repositories"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	expectedNum := 2
	got, err := repositories.SelectCommentList(testDB, articleID)

	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment {
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
			where comment_id = ?
		`
		testDB.Exec(sqlStr, newComment.CommentID)
	})
}
