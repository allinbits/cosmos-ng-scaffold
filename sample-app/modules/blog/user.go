package blog

import (
	"github.com/spf13/cobra"
	framework "github.com/tendermint/starport-framework/starport-framework"
)

func init() {
	framework.RegisterMessage(new(CreatePost), new(CreatePostHandler)) // associated with the Blog module with reflection by pkg name.
}

// CreatePost represents a create post request.
type CreatePost struct {
	Name    string
	Content string
}

// CreatePostHandler is a handler for CreatePost.
type CreatePostHandler struct{}

// Handle handles a new request for type.
func (h *CreatePostHandler) Handle(c *framework.ModuleContext, r framework.Request) error {
	var m CreatePost

	r.MustDecode(&m)

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
