package main

import (
	"github.com/joincloud/home-platform/service"
	financeCMD "github.com/joincloud/home-services/finance/cmd"
	"github.com/micro/go-micro/v2"
)

func main() {
	opts := &service.Options{
		Service: micro.NewService(micro.Name("home.srv.all")),
	}

	financeCMD.Init(opts)

	err := opts.Service.Run()
	if err != nil {
		panic(err)
	}
}
