package storage

import "github.com/pavva91/golang-api-structure/types"

type MongoStorage struct{}

func (s *MongoStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}
}
