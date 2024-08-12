package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/repo"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/emptypb" // Import the emptypb package
)

type CommentServiceServer struct {
	pb.UnimplementedCommentServiceServer
	Repo repo.CommentRepository
}

func NewCommentServiceServer(repo repo.CommentRepository) *CommentServiceServer {
	return &CommentServiceServer{
		Repo: repo,
	}
}

func (s *CommentServiceServer) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	comment := req.GetComment()

	// Validation
	if comment.CommentId == "" || comment.PostId == "" || comment.UserId == "" || comment.Content == "" {
		return nil, fmt.Errorf("commentId, postId, userId, and content are required")
	}

	createdAt := time.Now()
	updatedAt := time.Now()

	newComment := &models.Comment{
		CommentID: comment.CommentId,
		PostID:    comment.PostId,
		UserID:    comment.UserId,
		Content:   comment.Content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	err := s.Repo.Create(newComment)
	if err != nil {
		return nil, err
	}

	return &pb.CommentResponse{
		Comment: &pb.Comment{
			CommentId: newComment.CommentID,
			PostId:    newComment.PostID,
			UserId:    newComment.UserID,
			Content:   newComment.Content,
			CreatedAt: timestamppb.New(newComment.CreatedAt),
			UpdatedAt: timestamppb.New(newComment.UpdatedAt),
		},
	}, nil
}

func (s *CommentServiceServer) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.CommentResponse, error) {
	commentID := req.GetCommentId()
	postID := req.GetPostId()

	comment, err := s.Repo.GetByID(postID, commentID)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, fmt.Errorf("comment not found")
	}

	return &pb.CommentResponse{
		Comment: &pb.Comment{
			CommentId: comment.CommentID,
			PostId:    comment.PostID,
			UserId:    comment.UserID,
			Content:   comment.Content,
			CreatedAt: timestamppb.New(comment.CreatedAt),
			UpdatedAt: timestamppb.New(comment.UpdatedAt),
		},
	}, nil
}

func (s *CommentServiceServer) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error) {
	comment := req.GetComment()

	// Validation
	if comment.CommentId == "" || comment.PostId == "" || comment.UserId == "" || comment.Content == "" {
		return nil, fmt.Errorf("commentId, postId, userId, and content are required")
	}

	existingComment, err := s.Repo.GetByID(comment.PostId, comment.CommentId)
	if err != nil {
		return nil, err
	}
	if existingComment == nil {
		return nil, fmt.Errorf("comment not found")
	}
	if existingComment.UserID != comment.UserId {
		return nil, fmt.Errorf("user is not authorized to update this comment")
	}

	existingComment.Content = comment.Content
	existingComment.UpdatedAt = time.Now()

	err = s.Repo.Update(existingComment)
	if err != nil {
		return nil, err
	}

	return &pb.CommentResponse{
		Comment: &pb.Comment{
			CommentId: existingComment.CommentID,
			PostId:    existingComment.PostID,
			UserId:    existingComment.UserID,
			Content:   existingComment.Content,
			CreatedAt: timestamppb.New(existingComment.CreatedAt),
			UpdatedAt: timestamppb.New(existingComment.UpdatedAt),
		},
	}, nil
}

func (s *CommentServiceServer) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*emptypb.Empty, error) {
	commentID := req.GetCommentId()
	postID := req.GetPostId()

	existingComment, err := s.Repo.GetByID(postID, commentID)
	if err != nil {
		return nil, err
	}
	if existingComment == nil {
		return nil, fmt.Errorf("comment not found")
	}
	if existingComment.UserID != req.GetUserId() {
		return nil, fmt.Errorf("user is not authorized to delete this comment")
	}

	err = s.Repo.Delete(postID, commentID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *CommentServiceServer) ListComments(ctx context.Context, req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	postID := req.GetPostId()

	comments, err := s.Repo.ListByPostID(postID)
	if err != nil {
		return nil, err
	}

	var commentList []*pb.Comment
	for _, comment := range comments {
		commentList = append(commentList, &pb.Comment{
			CommentId: comment.CommentID,
			PostId:    comment.PostID,
			UserId:    comment.UserID,
			Content:   comment.Content,
			CreatedAt: timestamppb.New(comment.CreatedAt),
			UpdatedAt: timestamppb.New(comment.UpdatedAt),
		})
	}

	return &pb.ListCommentsResponse{Comments: commentList}, nil
}
