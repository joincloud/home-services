package node

import (
	"context"
	"sync"

	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
)

type Node struct {
	APIs map[string]API

	ctx context.Context
}

// One dir one API
type API struct {
	Name           string
	Dir            string
	IsTmp          bool
	BootstrapPeers []string
	ipfs           icore.CoreAPI
}

func (n *Node) connectToPeers() error {
	var wg sync.WaitGroup
	for _, api := range n.APIs {
		peerInfos := make(map[peer.ID]*peer.AddrInfo, len(api.BootstrapPeers))
		for _, bootsPeer := range api.BootstrapPeers {
			addr, err := ma.NewMultiaddr(bootsPeer)
			if err != nil {
				return err
			}
			pii, err := peer.AddrInfoFromP2pAddr(addr)
			if err != nil {
				return err
			}
			pi, ok := peerInfos[pii.ID]
			if !ok {
				pi = &peer.AddrInfo{ID: pii.ID}
				peerInfos[pi.ID] = pi
			}
			pi.Addrs = append(pi.Addrs, pii.Addrs...)
		}

		wg.Add(len(peerInfos))
		for _, peerInfo := range peerInfos {
			go func(peerInfo *peer.AddrInfo) {
				defer wg.Done()
				err := api.ipfs.Swarm().Connect(n.ctx, *peerInfo)
				if err != nil {
					log.Printf("failed to connect to %s: %s", peerInfo.ID, err)
				}
			}(peerInfo)
		}
		wg.Wait()
	}

	return nil
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
