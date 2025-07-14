package api

import (
	"context"
	"net/http"
)

type Controller interface {
	Download(ctx context.Context, w http.ResponseWriter, r *http.Request)
}
