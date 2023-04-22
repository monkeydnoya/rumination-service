package mongodriver

import (
	"time"

	"github.com/monkeydnoya/hiraishin-blog/pkg/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

type Blog struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" validate:"required"`
	Description string             `json:"description" validate:"required"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Author      domain.User        `json:"user" bson:"user"`
}

func toModel(blog Blog) domain.BlogResponse {
	return domain.BlogResponse{
		ID:          blog.ID.Hex(),
		Title:       blog.Title,
		Description: blog.Description,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
		Author:      blog.Author,
	}
}

// func toEntity(blog domain.BlogResponse) Blog {
// 	objID, err := primitive.ObjectIDFromHex(blog.ID)
// 	if err != nil {
// 		fmt.Printf("[ERROR] Can't convert string ID to object ID. Error: %s", err)
// 	}
// 	return Blog{
// 		ID:          objID,
// 		Title:       blog.Title,
// 		Description: blog.Description,
// 		CreatedAt:   blog.CreatedAt,
// 		UpdatedAt:   blog.UpdatedAt,
// 		Author:      blog.Author,
// 	}
// }
