package facade

import (
	"context"
	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/validation"
)

type PostFacade interface {
	CreatePost(ctx context.Context, post *pb.Post) (*pb.Post, error)
}

func NewPostFacade() PostFacade {
	return &PostFacadeImpl{}
}

type PostFacadeImpl struct {
}

func (p *PostFacadeImpl) CreatePost(ctx context.Context, post *pb.Post) (*pb.Post, error) {
	// create a post
	validation.ValidateCreatePost(post)
	return nil, nil
}
