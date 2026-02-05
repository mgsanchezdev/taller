package domain

import "time"

type MediaType string

const (
	MediaTypePhoto MediaType = "photo"
	MediaTypeVideo MediaType = "video"
	MediaTypePDF   MediaType = "pdf"
)

type Media struct {
	ID         string    `json:"id"`
	EntityType string    `json:"entity_type"` // entry, inspection, repair
	EntityID   string    `json:"entity_id"`
	Type       MediaType `json:"type"`
	URL        string    `json:"url"`
	Key        string    `json:"key"` // S3/storage key
	Filename   string    `json:"filename"`
	Size       int64     `json:"size"`
	CreatedAt  time.Time `json:"created_at"`
}
