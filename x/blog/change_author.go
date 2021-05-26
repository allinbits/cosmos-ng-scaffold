package blog

import (
	"github.com/allinbits/cosmos-ng-scaffold/x/blog/pb"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

func init() {
	initializers = append(initializers, func(client module.Client, builder *module.DescriptorBuilder) {
		builder.HandlesStateTransition(&pb.ChangeAuthor{}, &ChangeAuthor{client: &Client{Client: client}}, true)
	})
}

type ChangeAuthor struct {
	client *Client
}

func (c *ChangeAuthor) Exec(tr statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
	return
}
