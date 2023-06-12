package main

import (
	"log"
	"net/http"

	"andsp.id/andspdev/config"
	"andsp.id/andspdev/config/routes"
	"andsp.id/andspdev/config/variables"
)

func main() {
	var portServer = variables.PortServer

	config.ConnectDB()

	log.Println("Server berjalan pada port " + portServer + ". Dengan URL: http://localhost" + portServer)
	log.Fatal(http.ListenAndServe(portServer, routes.Routes()))
}
