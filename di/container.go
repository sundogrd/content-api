package di

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-api/grpc_gen/comment"
	"github.com/sundogrd/content-api/grpc_gen/content"
	"github.com/sundogrd/content-api/grpc_gen/user"
	commentGrpc "github.com/sundogrd/content-api/providers/grpc/comment"
	contentGrpc "github.com/sundogrd/content-api/providers/grpc/content"
	"github.com/sundogrd/content-api/utils/redis"
)

type Container struct {
	GormDB            *gorm.DB
	RedisClient       *redis.Client
	CommentGrpcClient comment.CommentServiceClient
	ContentGrpcClient content.ContentServiceClient
	UserGrpcClient    user.UserServiceClient
}

func InitContainer() (*Container, error) {
	commentClient, _, err := commentGrpc.NewGrpcCommentClient()
	if err != nil {
		logrus.Errorf("[Main] Init commentClient error: %+v", err)
		return nil, err
	}
	logrus.Infof("comment grpc client init success")

	contentClient, _, err := contentGrpc.NewGrpcContentClient()
	if err != nil {
		logrus.Errorf("[Main] Init commentClient error: %+v", err)
		return nil, err
	}
	logrus.Infof("content grpc client init success")

	return &Container{
		CommentGrpcClient: commentClient,
		ContentGrpcClient: contentClient,
	}, nil
}