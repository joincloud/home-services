package cmd

import (
	"github.com/joincloud/home-platform/service"
	"github.com/joincloud/home-services/finance/book"
	"github.com/joincloud/home-services/finance/book/handler"
	bProto "github.com/joincloud/home-services/proto/finance/book"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

func Init(opts *service.Options) {
	opts.Service.Init(micro.Action(func(context *cli.Context) error {
		book.Init()
		return nil
	}))

	err := bProto.RegisterBookHandler(opts.Service.Server(), new(handler.BookHandler))
	if err != nil {
		panic(err)
	}
}

func Run(opts *service.Options) {
	err := opts.Service.Run()
	if err != nil {
		panic(err)
	}
}
