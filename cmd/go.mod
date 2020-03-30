module github.com/joincloud/home-services/cmd

replace (
	github.com/joincloud/home-services/finance v1.0.0 => /Users/shuxian/Projects/joincloud/home-services/finance
	github.com/joincloud/home-services/proto v1.0.0 => /Users/shuxian/Projects/joincloud/home-services/proto
)

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/joincloud/home-services/finance v1.0.0
	github.com/joincloud/home-services/proto v1.0.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
)
