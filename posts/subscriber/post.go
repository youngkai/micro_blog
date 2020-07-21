package subscriber

import (
	"context"

	post "blog/posts/proto/posts"
)

type Post struct{}

func (e *Post) Handle(ctx context.Context, msg *post.Post) error {
	return nil
}

func Handler(ctx context.Context, msg *post.Post) error {
	return nil
}
