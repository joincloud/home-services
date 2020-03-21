module github.com/joincloud/home/finance

go 1.13

replace github.com/joincloud/home/proto v0.0.0 => /Users/shuxian/workspace/go/src/github.com/joincloud/home/proto

require (
	github.com/golang/protobuf v1.3.3
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/micro/go-micro/v2 v2.1.0 // indirect
	github.com/joincloud/home/proto v0.0.0
)
