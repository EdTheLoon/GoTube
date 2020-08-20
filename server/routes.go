package server

import "net/http"

func (s *Server) routes() {
	// Serves static asset files through a fileserver
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", s.fs))

	// Web Interface handling
	s.router.HandleFunc("/", s.viewHandler)

	// Get video details
	s.router.HandleFunc("/getVideoDetails/{url}", s.getVideoDetails)

	// Download video
	s.router.HandleFunc("/getVideo/{url}", s.getVideo)
}
