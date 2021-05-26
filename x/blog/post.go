package blog

import (
	"github.com/allinbits/cosmos-ng-scaffold/x/blog/pb"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/orm"
)

type Post = pb.Post

func init() {
	generatedInitializers = append(generatedInitializers, func(client module.Client) func(*module.DescriptorBuilder) {
		return func(builder *module.DescriptorBuilder) {
			builder.OwnsStateObject(&pb.Post{}, orm.RegisterOptions{PrimaryKey: "id"})
		}
	})
}
