package app

func NewAPI(s *Server) {
	s.Router.GET("/status", HeathCheck)
	s.Router.POST("/validate", ValidateUser)
}
