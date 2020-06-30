package main

import (
	"context"
	"github.com/joincloud/home-platform/home-services/conf"
	"github.com/joincloud/home-platform/home-services/node"
	"github.com/joincloud/home-platform/home-services/registry"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf.Init("/Users/shuxian/Projects/joincloud/home-services/application.yml")
	node.Init(ctx)
	registry.Register(ctx)
}
