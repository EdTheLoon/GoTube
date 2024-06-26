package server

import "net/http"

func (s *Server) routes() {
	// Serves static asset files through a fileserver
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", s.fs))

	// Web Interface handling
	s.router.HandleFunc("/gotube/", s.viewHandler)

	// Get video details
	s.router.HandleFunc("/gotube/getVideoDetails/{url}", s.getVideoDetails)

	// Download video
	s.router.HandleFunc("/gotube/getVideo/{url}/{format}/{filename}", s.getVideo)

	// Download audio
	s.router.HandleFunc("/gotube/getAudio/{url}/{format}/{filename}", s.getAudio)
}
