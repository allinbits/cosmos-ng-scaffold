package blog

import (
	"github.com/allinbits/cosmos-ng-scaffold/x/blog/pb"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

func init() {
	generatedInitializers = append(generatedInitializers, func(client module.Client) func(*module.DescriptorBuilder) {
		return func(builder *module.DescriptorBuilder) {
			builder.HandlesStateTransition(&pb.CreatePost{}, &CreatePost{client: &Client{Client: client}}, true)
		}
	})
}

type CreatePost struct {
	client *Client
}

func (c *CreatePost) Exec(tr statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
	create := tr.Transition.(*pb.CreatePost)
	post := Post{
		Id:      create.Id,
		Author:  tr.Users.List()[0].GetName(),
		Content: create.Content,
	}
	err = c.client.Create(&post)
	return
}
