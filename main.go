package main

import (
	"context"
	log "github.com/sirupsen/logrus"

	"github.com/joincloud/home-platform/home-services/conf"
	"github.com/joincloud/home-platform/home-services/node"
	"github.com/joincloud/home-platform/home-services/registry"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf.Init("/home/sx/Project/home/home-services/application.yml")
	node.Init(ctx)
	registry.Register(ctx)

	/*	err := node.GetWriteTo(ctx, "tmp", "QmUaoioqU7bxezBQZkUcgcSyokatMY71sxsALxQmRRrHrj", "/home/sx/Project/home/files/dir/hello")
		if err != nil {
			panic(err)
		}*/

	cID, err := node.ServeFile(ctx, "tmp", "/home/sx/Project/home/files/dir/helloSx.txt")
	if err != nil {
		panic(err)
	}

	log.Info(cID)

	select {
	// wait on context cancel
	case <-ctx.Done():
	}

	log.Info("over")
}
