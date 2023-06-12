package main

import (
	"log"
	"net/http"

	"andsp.id/andspdev/config"
)

func main() {
	var portServer = config.PortServer

	config.ConnectDB()
	
	log.Println("Server berjalan pada port "+portServer+". Dengan URL: http://localhost"+portServer)
	http.ListenAndServe(portServer, config.Routes())
}