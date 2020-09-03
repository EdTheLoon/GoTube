package server

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// The Server struct
// outputDir	 	The filesystem location of videos
// assetsDir		The filesystem location of static assets
// router			The HTTP router to be used
// fs				The FileServer for static files
// log				The location for log files
type Server struct {
	outputDir string
	assetsDir string
	router    *mux.Router
	fs        http.Handler
	log       *os.File
}

// NewServer creates a new server
func NewServer(outputDir string, assetsDir string, log string) Server {
	// Create/open the log file
	f, err := os.OpenFile(log, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	// Create the server struct
	s := Server{
		outputDir,
		assetsDir,
		mux.NewRouter(),
		http.FileServer(http.Dir(assetsDir)),
		f,
	}

	// Create the output directory if it does not exist
	_, err = os.Stat(s.outputDir)
	if err != nil {
		os.MkdirAll(s.outputDir, 0775)
		s.Log("Created output directory: " + s.outputDir)
	}

	// Initiate the routes
	s.routes()
	return s
}

// GetRouter returns the mux Router in use
func (s *Server) GetRouter() *mux.Router {
	return s.router
}

// GetDownloadsDir returns the directory to store downloads
func (s *Server) GetDownloadsDir() string {
	return s.outputDir
}
