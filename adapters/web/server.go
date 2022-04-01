package web

import (
	"github.com/labstack/echo"
	"log"
)

type Server struct {
	port   string
	engine *echo.Echo
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		engine: echo.New(),
	}
}

func (s *Server) Run() {
	router := UseRouter(s.engine)
	log.Fatal(router.Start(":" + s.port))
}
