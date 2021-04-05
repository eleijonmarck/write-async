package server

func (s *server) routes() {
	s.router.HandleFunc("/", s.HandleHealth())
	s.router.HandleFunc("/health", s.HandleHealth())
	s.router.HandleFunc("/job", s.HandleAddJob())
}
