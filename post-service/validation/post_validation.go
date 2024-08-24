package validation

import "github.com/nawarajshah/grpc-post-service/pb"

func ValidateCreatePost(post *pb.Post) {
	if post.Title == "" {
		panic("Title is required")

	}

	if len(post.Title) > 100 {
		// return error
		panic("Title is too long")
	}
}
