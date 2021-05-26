package blog

import (
	"github.com/spf13/cobra"
	framework "github.com/tendermint/starport-framework/starport-framework"
)

func init() {
	framework.RegisterMessage(new(CreatePost), new(CreatePostHandler)) // associated with the Blog module with reflection by pkg name.
}

func init() {
	framework.RegisterQuery(new(GetPostRequest), new(GetPostResponse), new(GetPostHandler))
}

// Post is a blog post.
type Post struct {
	ID      string
	Name    string
	Content string
}

// CreatePost represents a create post request.
type CreatePost struct {
	Post
}

// CreatePostHandler is a handler for CreatePost.
type CreatePostHandler struct {
	framework.Handler
}

// Handle handles a new request for type.
func (h *CreatePostHandler) Handle(c *framework.ModuleContext, m CreatePost) error {
	return c.Keeper.Save(m.Name, m)
}

// Command optionally add a command for this type.
func (h *CreatePostHandler) Command() *cobra.Command {
	return &cobra.Command{
		Use: "create-post",
	}
}

// Context called on each message handler once to optionnaly do advanced configuration on the
// module context.
func (h *CreatePostHandler) Context(c *framework.ModuleContext) {}

// GetPostRequest represents a get post request.
type GetPostRequest struct {
	ID string
}

// GetPostResponse represents a get post response payload.
type GetPostResponse struct {
	Post
}

// GetPostHandler is a get post handler.
type GetPostHandler struct {
	framework.Handler
}

// Handle handles get post query.
func (h *GetPostHandler) Handle(c *framework.ModuleContext, r GetPostRequest) (GetPostResponse, error) {
	var post Post

	if err := c.Keeper.Get(r.ID, &post); err != nil {
		return GetPostResponse{}, err
	}

	return GetPostResponse{Post: post}, nil
}
