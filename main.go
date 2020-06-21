package main

import (
	"github.com/joincloud/home-platform/home-services/conf"
	"github.com/joincloud/home-platform/home-services/registry"
)

func main() {
	conf.Init("E:\\GOPATH\\src\\github.com\\printfcoder\\home-services\\application.yml")
	registry.Register()
}
