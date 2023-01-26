package repositories

import (
	"database/sql"

	"github.com/an-0305/study-go-api/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now())
	`	

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.CommentID)
	if err != nil {
		return models.Comment{}, err
	}

	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	var commentArray []models.Comment

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}


	return commentArray, nil
}
