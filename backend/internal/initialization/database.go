package initialization

import (
	"context"
	"fmt"
	"time"

	"github.com/tienhai2808/anonymous_forest/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitDB(cfg *config.Config) (*mongo.Client, error) {
	mongoURI := cfg.Database.DBUri

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	clientOptions.SetMaxPoolSize(10)
	clientOptions.SetMinPoolSize(5)

	clientOptions.SetConnectTimeout(10 * time.Second)
	clientOptions.SetServerSelectionTimeout(5 * time.Second)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("kết nối MongoDB thất bại: %w", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping tới MongoDB thất bại: %w", err)
	}

	return client, nil
}
