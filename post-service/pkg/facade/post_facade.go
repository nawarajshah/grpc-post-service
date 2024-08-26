package facade

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/repo"
	"github.com/nawarajshah/grpc-post-service/post-service/validation"
)

type PostFacade interface {
	CreatePost(ctx context.Context, post *pb.Post) (*pb.Post, error)
	GetPost(ctx context.Context, postId string) (*pb.Post, error)
	UpdatePost(ctx context.Context, post *pb.Post) (*pb.Post, error)
	DeletePost(ctx context.Context, postId string, userId string) error
}

type PostFacadeImpl struct {
	postRepo repo.PostRepository
}

func NewPostFacade(postRepo repo.PostRepository) PostFacade {
	return &PostFacadeImpl{
		postRepo: postRepo,
	}
}

func (p *PostFacadeImpl) CreatePost(ctx context.Context, post *pb.Post) (*pb.Post, error) {
	// create a post
	validation.ValidateCreatePost(post)

	// Generate a unique post ID if one is not provided
	if post.PostId == "" {
		post.PostId = uuid.New().String()
	}

	// Convert pb.Post to models.Post for storage
	storedPost := &models.Post{
		PostID:      post.PostId,
		Title:       post.Title,
		Description: post.Description,
		CreatedBy:   post.UserId,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}

	// Save the post
	err := p.postRepo.Create(storedPost)
	if err != nil {
		return nil, fmt.Errorf("error creating post: %v", err)
	}

	return post, nil
}

func (p *PostFacadeImpl) GetPost(ctx context.Context, postId string) (*pb.Post, error) {
	post, err := p.postRepo.GetByID(postId)
	if err != nil {
		return nil, fmt.Errorf("error getting post: %v", err)
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	return &pb.Post{
		PostId:      post.PostID,
		Title:       post.Title,
		Description: post.Description,
		UserId:      post.CreatedBy,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}, nil
}

func (p *PostFacadeImpl) UpdatePost(ctx context.Context, post *pb.Post) (*pb.Post, error) {
	// Fetch the post to verify ownership
	storedPost, err := p.postRepo.GetByID(post.PostId)
	if err != nil {
		return nil, fmt.Errorf("error getting post: %v", err)
	}
	if storedPost == nil {
		return nil, fmt.Errorf("post not found")
	}

	// Check if the user attempting to update the post is the owner
	if storedPost.CreatedBy != post.UserId {
		return nil, fmt.Errorf("you do not have premission to update")
	}

	// Proceed with the update
	storedPost.Title = post.Title
	storedPost.Description = post.Description
	storedPost.UpdatedAt = post.UpdatedAt

	err = p.postRepo.Update(storedPost)
	if err != nil {
		return nil, fmt.Errorf("error updating post: %v", err)
	}

	return post, nil
}

func (p *PostFacadeImpl) DeletePost(ctx context.Context, postId string, userId string) error {
	// Fetch the post to verify ownership
	storedPost, err := p.postRepo.GetByID(postId)
	if err != nil {
		return fmt.Errorf("error getting post: %v", err)
	}
	if storedPost == nil {
		return fmt.Errorf("post not found")
	}

	// Check if the user attempting to delete the post is the owner
	if storedPost.CreatedBy != userId {
		return fmt.Errorf("you do not have premission to update")
	}

	// Proceed with the deletion
	err = p.postRepo.Delete(storedPost.PostID)

	if err != nil {
		return fmt.Errorf("error deleting post: %v", err)
	}

	return nil
}
