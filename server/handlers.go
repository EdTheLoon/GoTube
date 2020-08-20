package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) viewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	load := vars["page"] + ".html"
	if load == ".html" {
		load = "index.html"
	}
	p, err := s.loadPage(load)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	w.Write(p)
}

func (s *Server) getVideoDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]

	// CARRY OUT SOME URL VALIDATION HERE IN THE FUTURE
	s.Log("Getting video details: " + url)

}

func (s *Server) getVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]

	// CARRY OUT SOME URL VALIDATION HERE IN THE FUTURE
	s.Log("Downloading video: " + url)
}
