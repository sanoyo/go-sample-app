package persistence

import "github.com/jmoiron/sqlx"

type ContentRepository struct {
	DB *sqlx.DB
}

func NewContentRepository(db *sqlx.DB) *ContentRepository  {
	return &ContentRepository{
		DB: db,
	}
}

func (s *ContentRepository) FindContentByID(contentSummaryID string, language string) (*model.Content, error) {
	query := `
		SELECT
			content_id, title, description, content_type
		FROM
			contents
		WHERE
			content_id = ?
		AND
			language = ?
		;
	`

	noLanguageQuery := `
		SELECT
			content_id, title, content_url, description, content_type
		FROM
			contents
		WHERE
			content_summary_id = ?
		LIMIT
			1			
		;
	`

	var content model.Content
	err := s.DB.Get(&content, query, contentSummaryID, language)
	if err == nil {
		return &content, nil
	}

	// レコードがなかった場合、英語のコンテンツを検索する
	if err == sql.ErrNoRows {
		err = s.DB.Get(&content, query, contentSummaryID, "en")
		if err == nil {
			return &content, nil
		}
	}

	// レコードがなかった場合、先頭のデータを取得する
	if err == sql.ErrNoRows {
		err = s.DB.Get(&content, noLanguageQuery, contentSummaryID)
		if err == nil {
			return &content, nil
		}
	}

	return nil, err
}
