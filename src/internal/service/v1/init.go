package service_v1

import (
	"context"
	"fmt"
	"time"

	"github.com/golerplate/pkg/cache"
	database_v1 "github.com/golerplate/user-store-svc/internal/database/v1"
)

const (
	userCacheDuration = time.Hour * 24
)

func generateUserCacheKeyWithEmail(email string) string {
	return fmt.Sprintf("user-store-svc:user:email:%v", email)
}

func generateUserCacheKeyWithUserID(userID string) string {
	return fmt.Sprintf("user-store-svc:user:user_id:%v", userID)
}

func generateUserCacheKeyWithUsername(username string) string {
	return fmt.Sprintf("user-store-svc:user:username:%v", username)
}

func generateUserCacheKeyWithExternalID(externalID string) string {
	return fmt.Sprintf("user-store-svc:user:external_id:%v", externalID)
}

type service struct {
	store database_v1.Database
	cache cache.Cache
}

func NewUserStoreService(ctx context.Context, store database_v1.Database, cache cache.Cache) (*service, error) {
	return &service{
		store: store,
		cache: cache,
	}, nil
}
