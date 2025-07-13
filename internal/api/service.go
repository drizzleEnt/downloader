package api

var _ Server = (*server)(nil)

func New() Server {
	return &server{}
}

type server struct {
}
