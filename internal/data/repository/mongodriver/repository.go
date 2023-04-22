package mongodriver

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/monkeydnoya/hiraishin-blog/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b BlogDAO) GetBlogs() ([]domain.BlogResponse, error) {
	var blogList []Blog
	var blogListDomain []domain.BlogResponse

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := b.DB.Collection("Blog").Find(ctx, bson.M{})
	if err != nil {
		return blogListDomain, nil
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &blogList)
	if err != nil {
		return blogListDomain, nil
	}

	ListOfBlog := make([]domain.BlogResponse, len(blogList))
	for k, v := range blogList {
		ListOfBlog[k] = toModel(v)
	}
	return ListOfBlog, nil
}

func (b BlogDAO) GetBlogById(id string) (domain.BlogResponse, error) {
	var blog domain.BlogResponse

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.BlogResponse{}, err
	}

	err = b.DB.Collection("Blog").FindOne(ctx, bson.M{"_id": idObjectID}).Decode(&blog)
	if err != nil {
		return domain.BlogResponse{}, err
	}

	return blog, nil
}

func (b BlogDAO) CreateBlog(blog domain.Blog, user domain.User) (domain.BlogResponse, error) {
	blog.CreatedAt = time.Now()
	blog.Author = user

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := b.DB.Collection("Blog").InsertOne(ctx, &blog)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return domain.BlogResponse{}, errors.New("user with that email already exist")
		}
		return domain.BlogResponse{}, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}
	if _, err := b.DB.Collection("Blog").Indexes().CreateOne(ctx, index); err != nil {
		return domain.BlogResponse{}, errors.New("could not create index for email")
	}

	var newBlog Blog
	err = b.DB.Collection("Blog").FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newBlog)
	if err != nil {
		fmt.Println("[INFO] Not found")
		return domain.BlogResponse{}, err
	}

	return toModel(newBlog), nil
}

func (b BlogDAO) UpdateBlog(blog domain.BlogResponse, user domain.User) (domain.BlogResponse, error) {
	var blogInDB Blog
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	blog.Author.Email = user.Email
	blog.Author.Username = user.Username
	id, _ := primitive.ObjectIDFromHex(blog.ID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"title":       blog.Title,
		"description": blog.Description,
		"updated_at":  time.Now(),
		"author":      blog.Author,
		// Rethink: about edited user
	}}

	_, err := b.DB.Collection("Blog").UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.BlogResponse{}, err
	}

	err = b.DB.Collection("Blog").FindOne(ctx, bson.M{"_id": blog.ID}).Decode(&blogInDB)
	if err != nil {
		return domain.BlogResponse{}, err
	}

	return toModel(blogInDB), nil
}

func (b BlogDAO) DeleteBlog(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = b.DB.Collection("Blog").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}
