package server

import (
	"log"
	"net/http"
)


type Server struct {
	logger 		*log.Logger
	httpServer 	*http.Server
}