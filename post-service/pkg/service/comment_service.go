package service

import (
	"context"
	"fmt"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/utils"
	"time"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/repo"
)

type CommentServiceServer struct {
	pb.UnimplementedCommentServiceServer
	CommentRepo repo.CommentRepository
	PostRepo    repo.PostRepository
}

// NewCommentServiceServer is a constructor for CommentServiceServer
func NewCommentServiceServer(commentRepo repo.CommentRepository, postRepo repo.PostRepository) *CommentServiceServer {
	return &CommentServiceServer{
		CommentRepo: commentRepo,
		PostRepo:    postRepo,
	}
}

func (s *CommentServiceServer) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	post, err := s.PostRepo.GetByID(req.PostId)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	// Assuming the correct field name for the post owner's ID is 'UserID'
	isApproved := req.UserId == post.UserID

	comment := &models.Comment{
		CommentID:  utils.GenerateUUID(),
		PostID:     req.PostId,
		UserID:     req.UserId,
		Content:    req.Content,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
		IsApproved: isApproved,
		OwnerID:    post.UserID, // Store the post owner's ID in the comment
	}

	err = s.CommentRepo.Create(comment)
	if err != nil {
		return nil, err
	}

	return &pb.CommentResponse{
		CommentId:  comment.CommentID,
		PostId:     comment.PostID,
		UserId:     comment.UserID,
		Content:    comment.Content,
		CreatedAt:  comment.CreatedAt,
		UpdatedAt:  comment.UpdatedAt,
		IsApproved: comment.IsApproved,
	}, nil
}

func (s *CommentServiceServer) GetCommentByID(ctx context.Context, req *pb.GetCommentRequest) (*pb.CommentResponse, error) {
	comment, err := s.CommentRepo.GetByID(req.CommentId)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, fmt.Errorf("comment not found")
	}

	return &pb.CommentResponse{
		CommentId:  comment.CommentID,
		PostId:     comment.PostID,
		UserId:     comment.UserID,
		Content:    comment.Content,
		CreatedAt:  comment.CreatedAt, // Keep as int64 for response
		UpdatedAt:  comment.UpdatedAt, // Keep as int64 for response
		IsApproved: comment.IsApproved,
	}, nil
}

func (s *CommentServiceServer) GetCommentsByPostID(ctx context.Context, req *pb.GetCommentsByPostIDRequest) (*pb.GetCommentsByPostIDResponse, error) {
	comments, err := s.CommentRepo.GetByPostID(req.PostId)
	if err != nil {
		return nil, err
	}

	var pbComments []*pb.CommentResponse
	for _, comment := range comments {
		pbComments = append(pbComments, &pb.CommentResponse{
			CommentId:  comment.CommentID,
			PostId:     comment.PostID,
			UserId:     comment.UserID,
			Content:    comment.Content,
			CreatedAt:  comment.CreatedAt,
			UpdatedAt:  comment.UpdatedAt,
			IsApproved: comment.IsApproved,
		})
	}

	return &pb.GetCommentsByPostIDResponse{Comments: pbComments}, nil
}

func (s *CommentServiceServer) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error) {
	comment, err := s.CommentRepo.GetByID(req.CommentId)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, fmt.Errorf("comment not found")
	}

	comment.Content = req.Content
	comment.UpdatedAt = time.Now().Unix()

	err = s.CommentRepo.Update(comment)
	if err != nil {
		return nil, err
	}

	return &pb.CommentResponse{
		CommentId:  comment.CommentID,
		PostId:     comment.PostID,
		UserId:     comment.UserID,
		Content:    comment.Content,
		CreatedAt:  comment.CreatedAt,
		UpdatedAt:  comment.UpdatedAt,
		IsApproved: comment.IsApproved,
	}, nil
}

func (s *CommentServiceServer) ApproveComment(ctx context.Context, req *pb.ApproveCommentRequest) (*pb.CommentResponse, error) {
	comment, err := s.CommentRepo.GetByID(req.CommentId)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, fmt.Errorf("comment not found")
	}

	if comment.OwnerID != req.UserId {
		return nil, fmt.Errorf("only the post owner can approve comments")
	}

	comment.IsApproved = true
	comment.UpdatedAt = time.Now().Unix()

	err = s.CommentRepo.Update(comment)
	if err != nil {
		return nil, err
	}

	return &pb.CommentResponse{
		CommentId:  comment.CommentID,
		PostId:     comment.PostID,
		UserId:     comment.UserID,
		Content:    comment.Content,
		CreatedAt:  comment.CreatedAt,
		UpdatedAt:  comment.UpdatedAt,
		IsApproved: comment.IsApproved,
	}, nil
}

func (s *CommentServiceServer) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	err := s.CommentRepo.Delete(req.CommentId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCommentResponse{}, nil
}

//func mapToPBComment(comment *models.Comment) *pb.Comment {
//	return &pb.Comment{
//		CommentId:  comment.CommentID,
//		PostId:     comment.PostID,
//		UserId:     comment.UserID,
//		Content:    comment.Content,
//		IsApproved: comment.IsApproved,
//		CreatedAt:  comment.CreatedAt,
//	}
//}
