package main

import (
	"fmt"
	"net/http"

	"github.com/temirlanKabylbekov/calendar/pkg/logger"
)

type Server struct {
	Address string
	Log     *logger.Logger
}

func (s Server) Run() error {
	http.HandleFunc("/", s.handler)
	if err := http.ListenAndServe(s.Address, nil); err != nil {
		return err
	}
	return nil
}

func (s Server) handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	fmt.Fprintf(w, "Hello, your path is: %s!", path)
	s.Log.Infof("requested path: %s", path)
}
