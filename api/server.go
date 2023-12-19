package api

import (
	"encoding/json"
	"net/http"

	"github.com/pavva91/golang-api-structure/storage"
	"github.com/pavva91/golang-api-structure/util"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserByID)
	http.HandleFunc("/user/id", s.handleDeleteUserByID)
	http.HandleFunc("/foo", s.handleFoo)
	http.HandleFunc("/bar", s.handleBar)
	http.HandleFunc("/baz", s.handleBaz)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	_ = util.Round2Dec(10.3432)

	json.NewEncoder(w).Encode(user)

}

func (s *Server) handleDeleteUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	json.NewEncoder(w).Encode(user)

}
