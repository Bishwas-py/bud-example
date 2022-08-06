package controller

import (
	context "context"
	"github.com/matthewmueller/hackernews"
)

// Controller for posts
type Controller struct {
	HN *hackernews.Client
}

// Post struct
type Post struct {
	ID int `json:"id"`
}

// Index of posts
// GET
func (c *Controller) Index(ctx context.Context) (posts []*hackernews.Story, err error) {
	return c.HN.FrontPage(ctx)
}

// Show post
// GET :id
func (c *Controller) Show(ctx context.Context, id int) (post *Post, err error) {
	return &Post{
		ID: id,
	}, nil
}
