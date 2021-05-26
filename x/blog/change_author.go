package blog

import (
	"github.com/allinbits/cosmos-ng-scaffold/x/blog/pb"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

func init() {
	generatedInitializers = append(generatedInitializers, func(client module.Client) func(*module.DescriptorBuilder) {
		return func(builder *module.DescriptorBuilder) {
			builder.HandlesStateTransition(&pb.ChangeAuthor{}, &CreatePost{client: &Client{Client: client}}, true)
		}
	})
}

type ChangeAuthor struct {
	client *Client
}

func (c *ChangeAuthor) Exec(tr statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
	return
}
