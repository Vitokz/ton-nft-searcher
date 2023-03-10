package httpserver

func (s *Server) SetRoutes() {
	userG := s.server.Group("/user")
	userG.GET("/nft_items", s.UserNFTs)

	s.AddSwagger()
}
