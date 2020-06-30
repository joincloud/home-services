package node

import "context"

func Init(ctx context.Context) {
	createRepos(ctx)
}
