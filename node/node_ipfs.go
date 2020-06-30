package node

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	config "github.com/ipfs/go-ipfs-config"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/home-platform/home-services/conf"
	log "github.com/sirupsen/logrus"
)

func createRepos(ctx context.Context) {
	if err := setupPlugins(""); err != nil {
		// todo error
		panic(err)
	}

	files := conf.Configs.Home.Files
	for s, f := range files {
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

		_, err = createNode(ctx, repoPath)
		if err != nil {
			// todo error
			panic(err)
		}
	}
}

func createNode(ctx context.Context, repoPath string) (icore.CoreAPI, error) {
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

func setupPlugins(externalPluginsPath string) error {
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
