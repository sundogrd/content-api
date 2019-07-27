package di

import (
	"github.com/jinzhu/gorm"
	"github.com/sundogrd/content-api/grpc_gen/comment"
	"github.com/sundogrd/content-api/grpc_gen/content"
	"github.com/sundogrd/content-api/grpc_gen/user"
	"github.com/sundogrd/content-api/utils/redis"
)

type Container struct {
	GormDB            *gorm.DB
	RedisClient       *redis.Client
	CommentGrpcClient comment.CommentServiceClient
	ContentGrpcClient content.ContentServiceClient
	UserGrpcClient    user.UserServiceClient
}
