package node

import (
	"context"
	"fmt"

	files "github.com/ipfs/go-ipfs-files"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	log "github.com/sirupsen/logrus"
)

func GetWriteTo(ctx context.Context, apiName, fileID, outPath string) (err error) {
	if api, ok := node.APIs[apiName]; ok {
		cID := icorepath.New(fileID)
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
