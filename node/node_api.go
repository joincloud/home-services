package node

import (
	"context"
	"fmt"

	files "github.com/ipfs/go-ipfs-files"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
)

func GetNodeAddrs(ctx context.Context, apiName string) (addrs []ma.Multiaddr, err error) {
	if api, ok := node.APIs[apiName]; ok {
		addrs, err := api.ipfs.Swarm().LocalAddrs(ctx)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return addrs, nil
	} else {
		return nil, fmt.Errorf("api %s is nonexisted", apiName)
	}
}

func GetWriteTo(ctx context.Context, apiName, fileID, outPath string) (err error) {
	if api, ok := node.APIs[apiName]; ok {
		cID := icorepath.New(fileID)
		log.Infof("try to get file %s from %s", fileID, cID)

		fileNode, err := api.ipfs.Unixfs().Get(ctx, cID)
		if err != nil {
			log.Error(err)
			return err
		}

		err = files.WriteTo(fileNode, outPath)
		if err != nil {
			log.Error(err)
			return err
		}

		return nil
	} else {
		return fmt.Errorf("api %s is nonexisted", apiName)
	}
}
