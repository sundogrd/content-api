package test

import (
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-api/di"
	"github.com/sundogrd/content-api/providers/grpc/comment"
	"github.com/sundogrd/content-api/providers/grpc/content"
)

func InitTestContainer() (*di.Container, error) {
	commentClient, _, err := comment.NewGrpcCommentClient()
	if err != nil {
		logrus.Errorf("[Main] Init commentClient error: %+v", err)
		return nil, err
	}
	logrus.Infof("comment grpc client init success")

	contentClient, _, err := content.NewGrpcContentClient()
	if err != nil {
		logrus.Errorf("[Main] Init commentClient error: %+v", err)
		return nil, err
	}
	logrus.Infof("content grpc client init success")

	return &di.Container{
		CommentGrpcClient: commentClient,
		ContentGrpcClient: contentClient,
	}, nil
}