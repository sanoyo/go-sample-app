package repository

import (
	"github.com/go-sample-app/app/domain/model"
)

type ContentAccesor interface {
	FindContentByID(contantSummaryID string) (*model.Content, error)
}
