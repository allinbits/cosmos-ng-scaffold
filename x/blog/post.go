package blog

import (
	"github.com/allinbits/cosmos-ng-scaffold/x/blog/pb"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/orm"
)

type Post = pb.Post

func init() {
	initializers = append(initializers, func(client module.Client, builder *module.DescriptorBuilder) {
		builder.OwnsStateObject(&pb.Post{}, orm.RegisterOptions{PrimaryKey: "id"})
	})
}
