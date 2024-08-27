package endpoint

import (
	"context"
	"fmt"
	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/facade"
)

type PostServiceRpcImpl struct {
	pb.UnimplementedPostServiceServer
	postFacade facade.PostFacade
}

func NewPostServiceRpc(postFacade facade.PostFacade) pb.PostServiceServer {
	return &PostServiceRpcImpl{
		postFacade: postFacade,
	}
}

func (s *PostServiceRpcImpl) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()
	post, err := s.postFacade.CreatePost(ctx, req.Post)

	if err != nil {
		return nil, err
	}
	return &pb.PostResponse{
		Post: post,
	}, nil
}

func (s *PostServiceRpcImpl) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	post, err := s.postFacade.GetPost(ctx, req.PostId)
	if err != nil {
		return nil, err
	}

	return &pb.PostResponse{
		Post: post,
	}, nil
}

func (s *PostServiceRpcImpl) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	post, err := s.postFacade.UpdatePost(ctx, req.Post)
	if err != nil {
		return nil, err
	}

	return &pb.PostResponse{
		Post: post,
	}, nil
}

func (s *PostServiceRpcImpl) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	err := s.postFacade.DeletePost(ctx, req.PostId, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePostResponse{
		Status: "Post deleted successfully",
		UserId: req.UserId,
	}, nil
}
