package node

import (
	"context"

	icore "github.com/ipfs/interface-go-ipfs-core"
)

type Node struct {
	APIs map[string]API
}

// One dir one node
type API struct {
	Name string
	Dir  string
	ipfs icore.CoreAPI
}

func (n *Node) connectToPeers() {

}

func Init(ctx context.Context) {
	if err := initPlugins(ctx, ""); err != nil {
		// todo error
		panic(err)
	}

	if err := initNode(ctx); err != nil {
		// todo error
		panic(err)
	}
}
