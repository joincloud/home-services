package book

import (
	"github.com/joincloud/home-services/finance/book/repository"
	_ "github.com/joincloud/home-services/finance/book/repository/sqlite"
)

func Init() {
	repository.Init()
}
