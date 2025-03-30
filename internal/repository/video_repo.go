package repository

import (
	"context"
	"fampay_youtube_fetcher/internal/pojo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoRepo struct {
	collection *mongo.Collection
}

func NewVideoRepo(db *mongo.Database) *VideoRepo {
	collection := db.Collection("videos")

	// Create indexes
	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "published_at", Value: -1}},
			Options: options.Index().SetName("published_at_idx"),
		},
		{
			Keys:    bson.D{{Key: "video_id", Value: -1}},
			Options: options.Index().SetName("video_id_idx").SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "title", Value: "text"}, {Key: "description", Value: "text"}},
			Options: options.Index().SetName("text_search_idx"),
		},
	})
	if err != nil {
		panic(err)
	}

	return &VideoRepo{collection: collection}
}

func (r *VideoRepo) Create(video *pojo.VideoMetaData) error {
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

func (r *VideoRepo) GetLatestPublishedDate() (*time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Find().
		SetSort(bson.D{{Key: "published_at", Value: -1}}).
		SetLimit(int64(1))

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var videos []pojo.VideoMetaData
	if err = cursor.All(ctx, &videos); err != nil {
		return nil, err
	}

	if len(videos) == 0 {
		return nil, nil
	}

	return &videos[0].PublishedAt, nil
}
