package server

func (s *Server) routes(){
	// TODO: add cors,subrouter,auth in future
	s.Router.Get("/employees", s.getEmployees())
	s.Router.Get("/start", s.getStartSoon())
}