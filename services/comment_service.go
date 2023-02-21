package services

import (
	"github.com/an-0305/study-go-api/apperrors"
	"github.com/an-0305/study-go-api/models"
	"github.com/an-0305/study-go-api/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
