package blog

import (
	"github.com/spf13/cobra"
	framework "github.com/tendermint/starport-framework/starport-framework"
)

func init() {
	framework.RegisterMessage(&CreatePost{}) // associated with the Blog module with reflection by pkg name.
}

// CreatePost represents a create post request.
type CreatePost struct {
	Name    string
	Content string
}

// Context sets type related configuration to the module context.
func (p *CreatePost) Context(c *framework.ModuleContext) {
	c.Command.AddCommand(p.createPostCommand())
}

// Handle handles a new request for type.
func (p *CreatePost) Handle(f framework.Request) error {
	var m CreatePost
	if err := f.Decode(&m); err != nil {
		return err
	}

	return f.Keeper().Save(m.Name, m)
}

func (p *CreatePost) createPostCommand() *cobra.Command {
	return &cobra.Command{
		Use: "create-post",
	}
}
