package main

import (
	"github.com/joincloud/home-platform/service"
	"github.com/joincloud/home-services/finance/cmd"
	"github.com/micro/go-micro/v2"
)

func main() {
	opts := &service.Options{
		Service: micro.NewService(micro.Name("home.srv.fin")),
	}

	cmd.Init(opts)
	cmd.Run(opts)
}
