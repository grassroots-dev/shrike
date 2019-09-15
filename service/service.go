package service

// Server is a type represneting a configured server.
type Server struct {
	DB      string
	Cache   string
	Storage string
}

// NewService returns a new service configured with all the resources.
func NewService(db string, cache string, storage string) *Server {
	return &Server{DB: "hello db", Cache: "hello cache", Storage: "hello storage"}
}
