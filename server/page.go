package server

import (
	"io/ioutil"
)

func (s *Server) loadPage(p string) ([]byte, error) {
	s.Log("Loading page: " + s.assetsDir + p)
	page, err := ioutil.ReadFile(s.assetsDir + p)
	return page, err
}
