package main

import (
	"context"

	"github.com/joincloud/home-platform/home-services/conf"
	"github.com/joincloud/home-platform/home-services/node"
	"github.com/joincloud/home-platform/home-services/registry"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf.Init("/home/sx/Project/home/home-services/application.yml")
	node.Init(ctx)
	registry.Register(ctx)

	addrs, err := node.GetNodeAddrs(ctx, "tmp")
	if err != nil {
		log.Error(err)
	}

	log.Info(addrs)
}
