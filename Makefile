GOPATH ?= $(HOME)/go

.PHONY: protoc
protoc:
	protoc --go_out=paths=source_relative:. x/blog/pb/post.proto x/blog/pb/create_post.proto