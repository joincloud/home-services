package node

import (
	"context"
	"io/ioutil"
	"sync"

	config "github.com/ipfs/go-ipfs-config"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/home-platform/home-services/conf"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
)

var (
	node *Node
	mux  sync.Mutex
)

type Node struct {
	APIs map[string]*API

	ctx context.Context
}

// One dir one API
type API struct {
	Name           string
	Dir            string
	IsTmp          bool
	BootstrapPeers map[string]string
	ipfs           icore.CoreAPI
}

func (n *Node) init() (err error) {
	log.Info("init node")
	err = n.prepareAPIs()
	if err != nil {
		return err
	}

	err = n.connectToPeers()
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) prepareAPIs() error {
	log.Info("init node, prepare apis")
	fs := conf.Configs.Home.Files
	for k, f := range fs {
		log.Infof("create dir: %s: %s", k, f.Dir)
		if f.IsTemp {
			go func() {
				select {
				case <-n.ctx.Done():
					// todo delete dir
				}
			}()
		}

		repoPath, err := ioutil.TempDir(f.Dir, "ipfs-shell")
		if err != nil {
			// todo error
			panic(err)
		}

		// Create a config with default options and a 2048 bit key
		cfg, err := config.Init(ioutil.Discard, 2048)
		if err != nil {
			// todo error
			panic(err)
		}

		err = fsrepo.Init(repoPath, cfg)
		if err != nil {
			// todo error
			panic(err)
		}

		ipfs, err := prepareAPI(n.ctx, repoPath)
		if err != nil {
			// todo error
			panic(err)
		}

		node.APIs[k] = &API{
			Name:           k,
			Dir:            f.Dir,
			IsTmp:          f.IsTemp,
			BootstrapPeers: f.BootstrapNodes,
			ipfs:           ipfs,
		}
	}

	return nil
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
					log.Errorf("failed to connect to %s: %s", peerInfo.ID, err)
				}
			}(peerInfo)
		}
		wg.Wait()
	}

	return nil
}

func prepareNode(ctx context.Context) {
	node = &Node{
		ctx:  ctx,
		APIs: map[string]*API{},
	}

	if err := node.init(); err != nil {
		// todo error
		panic(err)
	}
}

func Init(ctx context.Context) {
	mux.Lock()
	defer mux.Unlock()

	if err := initPlugins(ctx, ""); err != nil {
		// todo error
		panic(err)
	}

	prepareNode(ctx)
}
