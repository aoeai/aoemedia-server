package user

import (
	"github.com/aoemedia-server/config"
	"github.com/aoemedia-server/domain/user"
	"sync"
)

type Repository struct {
	repository *user.Repository
}

var (
	instance *Repository
	once     sync.Once
	users    map[string]int64
)

func Inst() *Repository {
	once.Do(func() {
		instance = &Repository{}
	})
	initUsers()
	return instance
}

func initUsers() {
	users = make(map[string]int64)
	for _, user := range config.Inst().Users {
		users[user.Token] = user.ID
	}
}

func (r *Repository) GetIdByToken(token string) int64 {
	if id, ok := users[token]; ok {
		return id
	}
	return 0
}
