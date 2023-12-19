package storage

import "github.com/pavva91/golang-api-structure/types"

type MemoryStorage struct {
	
}

// constructor
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}
}
