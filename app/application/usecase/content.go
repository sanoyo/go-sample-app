package usecase

import (
	"github.com/go-sample-app/app/domain/model"
	"github.com/go-sample-app/app/domain/service"
	"github.com/go-sample-app/app/infrastructure"
	"github.com/go-sample-app/app/infrastructure/persistence"
)

func ShowContent(contentID string) (*model.Content, error) {
	contentRepositoryAccessor := persistence.NewContentRepository(infrastructure.DB)
	contentServiceAccesor := service.NewContentService(contentRepositoryAccessor)

	content, err := contentServiceAccesor.ShowContent(contentID)
	if err != nil {
		return nil, err
	}

	return content, nil
}
