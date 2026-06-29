package api

import (
	"photoshop-like/internal/engine"
	"photoshop-like/internal/session"
)

type Server struct {
	engine   *engine.Engine
	session  *session.Store
	openvino *engine.OpenVino
}

func New(eng *engine.Engine, openvino *engine.OpenVino, store *session.Store) *Server {
	return &Server{eng, store, openvino}
}
