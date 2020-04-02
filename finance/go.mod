module github.com/joincloud/home-services/finance

go 1.13

replace (
   github.com/joincloud/home-services/proto v1.0.0 => /home/sx/Project/go/joincloud/home-services/proto
"github.com/joincloud/home-platform" v1.0.0 =>  /home/sx/Project/go/joincloud/home-platform
)

require (
	github.com/golang/protobuf v1.3.5
	github.com/joincloud/home-services/proto v1.0.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.3.0
	github.com/joincloud/home-platform v1.0.0
)
