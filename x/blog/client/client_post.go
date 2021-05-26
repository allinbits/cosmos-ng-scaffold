package client

import (
	"github.com/allinbits/cosmos-ng-scaffold/x/blog/pb"
	"github.com/fdymylja/tmos/runtime/meta"
)

func (c *Client) GetPost(id string) (*pb.Post, error) {
	var post pb.Post
	return &post, c.Client.Get(meta.NewStringID(id), &post)
}
