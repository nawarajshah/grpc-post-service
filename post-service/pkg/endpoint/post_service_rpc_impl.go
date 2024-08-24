package service

import (
	"context"
	"fmt"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/facade"
	"time"

	"github.com/nawarajshah/grpc-post-service/pb"
)

type PostServiceRpcImpl struct {
	pb.UnimplementedPostServiceServer
	postFacade facade.PostFacade
}

// NewPostServiceServer is a constructor for GrpcService
func NewPostServiceRpc(postFacade facade.PostFacade) pb.PostServiceServer {
	return PostServiceRpcImpl{
		postFacade: postFacade,
	}
}

func (s *PostServiceRpcImpl) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	// Generate a unique post ID if one is not provided
	//postID := req.Post.PostId

	defer func() {
		if err := recover(); err != nil {
			// return error
			//
		}
	}()
	post, err := s.postFacade.CreatePost(ctx, req.Post)
	//if postID == "" {
	//postID := uuid.New().String()
	////}
	//
	//post := &models.Post{
	//	PostID:      postID,
	//	Title:       req.Post.Title,
	//	Description: req.Post.Description,
	//	CreatedBy:   req.Post.UserId,
	//	CreatedAt:   time.Now().Unix(),
	//	UpdatedAt:   time.Now().Unix(),
	//}
	//
	//err := s.PostRepo.Create(post)
	//if err != nil {
	//	return nil, fmt.Errorf("error creating post: %v", err)
	//}

	if err != nil {
		//return nil, err
	}
	return &pb.PostResponse{
		Post: post,
		//Post: &pb.Post{
		//	PostId:      post.PostID,
		//	Title:       post.Title,
		//	Description: post.Description,
		//	UserId:      post.CreatedBy,
		//	CreatedAt:   post.CreatedAt,
		//	UpdatedAt:   post.UpdatedAt,
		//},
	}, nil
}

func (s *GrpcService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	post, err := s.PostRepo.GetByID(req.PostId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving post: %v", err)
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:      post.PostID,
			Title:       post.Title,
			Description: post.Description,
			UserId:      post.CreatedBy, // Update to CreatedBy
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
		},
	}, nil
}

func (s *GrpcService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	// Fetch the post to verify ownership
	post, err := s.PostRepo.GetByID(req.Post.PostId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving post: %v", err)
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	// Check if the user attempting to update the post is the owner
	if post.CreatedBy != req.Post.UserId {
		return nil, fmt.Errorf("you do not have permission to update this post")
	}

	// Proceed with the update
	post.Title = req.Post.Title
	post.Description = req.Post.Description
	post.UpdatedAt = time.Now().Unix()

	err = s.PostRepo.Update(post)
	if err != nil {
		return nil, fmt.Errorf("error updating post: %v", err)
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:      post.PostID,
			Title:       post.Title,
			Description: post.Description,
			UserId:      post.CreatedBy,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
		},
	}, nil
}

func (s *GrpcService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	// Fetch the post to verify ownership
	post, err := s.PostRepo.GetByID(req.PostId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving post: %v", err)
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	// Check if the user attempting to delete the post is the owner
	if post.CreatedBy != req.UserId {
		return nil, fmt.Errorf("you do not have permission to delete this post")
	}

	// Proceed with the deletion
	err = s.PostRepo.Delete(req.PostId)
	if err != nil {
		return nil, fmt.Errorf("error deleting post: %v", err)
	}

	return &pb.DeletePostResponse{
		Status: "Post deleted successfully",
		UserId: post.CreatedBy,
	}, nil
}
