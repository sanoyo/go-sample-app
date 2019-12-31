package service

import (
	"github.com/go-sample-app/app/domain/model"
	"github.com/go-sample-app/app/domain/repository"
)

type ContentService struct {
	contentAccesor repository.ContentAccesor
}

func NewContentService(c repository.ContentAccesor) *ContentService {
	return &ContentService{
		contentAccesor: c,
	}
}

func (c *ContentService) ShowContent(contentID string) (*model.Content, error) {
	content, err := c.contentAccessor.FindContentByID(contentID)
	if err != nil {
		return nil, err
	}

	return content, nil
}
