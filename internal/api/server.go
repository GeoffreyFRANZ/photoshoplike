package api

import (
	"photoshop-like/internal/engine"
	"photoshop-like/internal/session"
)

type Server struct {
	engine  *engine.Engine
	session *session.Store
}

func New(eng *engine.Engine, store *session.Store) *Server {
	return &Server{eng, store}
}
