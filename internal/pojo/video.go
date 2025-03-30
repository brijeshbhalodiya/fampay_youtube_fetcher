package pojo

import "time"

type VideoMetaData struct {
	Id               string    `json:"id" bson:"_id,omitempty"`
	VideoId          string    `json:"video_id" bson:"video_id"`
	Title            string    `json:"title" bson:"title"`
	Description      string    `json:"description" bson:"description"`
	PublishedAt      time.Time `json:"published_at" bson:"published_at"`
	DefaultThumbnail string
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}
