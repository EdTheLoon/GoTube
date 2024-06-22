package server

import (
	"net/http"
	"net/url"

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
		s.Log("Page not found: " + load)
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
	url := s.decode(vars["url"])
	format := vars["format"]
	filename := s.decode(vars["filename"])

	// CARRY OUT SOME URL VALIDATION HERE IN THE FUTURE
	s.Log("Downloading video (" + format + "): " + url)
	downloadVideo(url, format, s, filename)

	// For now just redirect to the folder for download
	http.Redirect(w, r, "https://s.reids.scot/yt/", http.StatusSeeOther)

	// TO DO: Send the response HTML pages.
}

func (s *Server) getAudio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := s.decode(vars["url"])
	format := vars["format"]
	filename := s.decode(vars["filename"])

	// CARRY OUT SOME URL VALIDATION HERE IN THE FUTURE
	s.Log("Downloading audio (" + format + "): " + url)
	downloadAudio(url, format, s, filename)

	// For now just redirect to the folder for download
	http.Redirect(w, r, "https://s.reids.scot/yt/", http.StatusSeeOther)

	// TO DO: Send the response HTML pages.
}

func (s *Server) decode(u string) (string) {
	d,_ := url.QueryUnescape(u)
	ds := string(d)
	s.Log("Encoded: " + u + "\nDecoded: " + ds)
	return ds
}
