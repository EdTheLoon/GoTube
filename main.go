package main

import (
	"net/http"

	"github.com/EdTheLoon/GoTube/server"
)

func main() {
	// Create a new server
	s := server.NewServer("/tmp", "/srv/gotube/assets", "/var/log/gotube.log")
	defer s.CloseLog()

	// TODO: Get specified port from config or command line
	port := "6060"

	// Start the web server and listen on specified port
	s.Log("Server starting on http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, s.GetRouter()); err != nil {
		panic(err)
	}
}
