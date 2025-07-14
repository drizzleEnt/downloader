package api

import (
	"context"
	"net/http"
)

var _ Controller = (*server)(nil)

func New() Controller {
	return &server{}
}

type server struct {
}

// Download implements Controller.
func (s *server) Download(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("download"))
	//w.Write([]byte("download"))
}
