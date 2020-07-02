package node

import (
	"context"

	files "github.com/ipfs/go-ipfs-files"
)

func GetWriteTo(ctx context.Context, apiName, fileID, path string) (err error) {
	if api, ok := node.APIs[apiName]; ok {
		cidFile, err := api.ipfs.Unixfs().Add(ctx, someFile)
		rootNodeFile, err := api.ipfs.Unixfs().Get(ctx, cidFile)
	}

	files.WriteTo(node.APIs[apiName].ipfs, path)
}
