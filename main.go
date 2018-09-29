package main

import (
	"log"
	"os"

	"github.com/go-chi/chi"

	"github.com/BottleneckStudio/go-web-starter-kit/homepage"
	"github.com/BottleneckStudio/go-web-starter-kit/server"
)

func main() {
	//Declare Config
	logger := log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lshortfile)
	router := chi.NewMux()
	serverAddress := ":1437"

	//Create new Server
	serverConfig := server.Config{
		Server: &server.Server{
			Logger:        logger,
			Router:        router,
			ServerAddress: serverAddress,
		},
	}
	srv := server.New(&serverConfig)

	//Setup Domain specific routes
	h := homepage.New(srv.Logger)
	h.SetupRoutes(srv.Router)

	//Start Server
	srv.Start()
}
