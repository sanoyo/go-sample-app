package model

// Contentがどういうものかを表す
type Content struct {
	ContentID   uint   `json:"content_id" db:"content_id"`
	Title       string `json:"title" db:"title"`
	ContentURL  string `json:"content_url" db:"content_url"`
	Description string `json:"description" db:"description"`
	ContentType uint   `json:"content_type" db:"content_type"`
	Language    uint   `json:"language" db:"language"`
}
