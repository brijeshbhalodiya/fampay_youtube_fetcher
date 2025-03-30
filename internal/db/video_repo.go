package db

import (
	"context"
	"fampay_youtube_fetcher/internal/pojo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type videoRepo struct {
	collection *mongo.Collection
}

func NewVideoRepo(db *mongo.Database) *videoRepo {
	collection := db.Collection("videos")

	// Create indexes
	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "published_at", Value: -1}},
			Options: options.Index().SetName("published_at_idx"),
		},
		{
			Keys:    bson.D{{Key: "title", Value: "text"}, {Key: "description", Value: "text"}},
			Options: options.Index().SetName("text_search_idx"),
		},
	})
	if err != nil {
		panic(err)
	}

	return &videoRepo{collection: collection}
}

func (r *videoRepo) Create(video *pojo.VideoMetaData) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	video.CreatedAt = time.Now()
	video.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, video)
	if err != nil {
		return err
	}

	return nil
}
