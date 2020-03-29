module github.com/joincloud/home-services/finance

go 1.13

replace github.com/joincloud/home-services/proto v1.0.0 => /Users/shuxian/Projects/joincloud/home-services/proto

require (
	github.com/golang/protobuf v1.3.5
	github.com/joincloud/home-services/proto v1.0.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
)
