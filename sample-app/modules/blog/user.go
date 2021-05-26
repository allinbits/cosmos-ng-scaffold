package blog

import framework "github.com/tendermint/starport-framework/starport-framework"

func init() {
	framework.RegisterMessage(&CreatePost{})
}

// CreatePost represents a create post request.
type CreatePost struct {
	Name    string
	Content string
}

// Handle handles a new request for type.
func (p *CreatePost) Handle(f framework.Request) error {
	var m CreatePost
	if err := f.Decode(&m); err != nil {
		return err
	}

	return f.Keeper().Save(m.Name, m)
}
