package storage

import "github.com/pavva91/golang-api-structure/types"

type Storage interface {
	Get(int) *types.User
}
