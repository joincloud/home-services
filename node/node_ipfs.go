package node

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	config "github.com/ipfs/go-ipfs-config"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/joincloud/home-platform/home-services/conf"
	"github.com/libp2p/go-libp2p-core/peer"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
)

var (
	coreAPI  icore.CoreAPI
	nodesMap map[string]Node = map[string]Node{}
)

func initPlugins(ctx context.Context, externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

func initNode(ctx context.Context) error {
	fs := conf.Configs.Home.Files
	for s, f := range fs {
		log.Infof("create dir: %s: %s", s, f.Dir)
		if f.IsTemp {
			go func() {
				select {
				case <-ctx.Done():
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

		// begin the node
		ipfs, err := prepareAPI(ctx, repoPath)
		if err != nil {
			// todo error
			panic(err)
		}

		// connectToPeers
		go connectToPeers(ctx, ipfs, conf.GetBootStrapNodes())

		exampleCIDStr := "QmUaoioqU7bxezBQZkUcgcSyokatMY71sxsALxQmRRrHrj"
		outputPath := "/Users/shuxian/Projects/joincloud/files/tmp/" + exampleCIDStr
		testCID := icorepath.New(exampleCIDStr)

		rootNode, err := ipfs.Unixfs().Get(ctx, testCID)
		if err != nil {
			panic(fmt.Errorf("Could not get file with CID: %s", err))
		}

		err = files.WriteTo(rootNode, outputPath)
		if err != nil {
			panic(fmt.Errorf("Could not write out the fetched CID: %s", err))
		}
	}

	return nil
}

func prepareAPI(ctx context.Context, repoPath string) (icore.CoreAPI, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node

	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}

	// Attach the Core API to the constructed node
	return coreapi.NewCoreAPI(node)
}

func getUnixfsNode(path string) (files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func connectToPeers(ctx context.Context, ipfs icore.CoreAPI, peers []string) error {
	var wg sync.WaitGroup
	peerInfos := make(map[peer.ID]*peerstore.PeerInfo, len(peers))
	for _, addrStr := range peers {
		addr, err := ma.NewMultiaddr(addrStr)
		if err != nil {
			return err
		}
		pii, err := peerstore.InfoFromP2pAddr(addr)
		if err != nil {
			return err
		}
		pi, ok := peerInfos[pii.ID]
		if !ok {
			pi = &peerstore.PeerInfo{ID: pii.ID}
			peerInfos[pi.ID] = pi
		}
		pi.Addrs = append(pi.Addrs, pii.Addrs...)
	}

	wg.Add(len(peerInfos))
	for _, peerInfo := range peerInfos {
		go func(peerInfo *peerstore.PeerInfo) {
			defer wg.Done()
			err := ipfs.Swarm().Connect(ctx, *peerInfo)
			if err != nil {
				log.Printf("failed to connect to %s: %s", peerInfo.ID, err)
			}
		}(peerInfo)
	}
	wg.Wait()
	return nil
}
