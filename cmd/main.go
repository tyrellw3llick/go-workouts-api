package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/tyrellw3llick/femProject/internal/app"
	"github.com/tyrellw3llick/femProject/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Application started on port: %d", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
