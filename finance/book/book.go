package book

import (
	"github.com/joincloud/home/finance/book/repository"
	_ "github.com/joincloud/home/finance/book/repository/sqlite"
)

func Init() {
	repository.Init()
}
