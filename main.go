package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/Dev-ManavSethi/my-website/models"

	"golang.org/x/net/websocket"

	"github.com/Dev-ManavSethi/my-website/controllers"
	"github.com/Dev-ManavSethi/my-website/utils"
	"github.com/joho/godotenv"
)

var tpl *template.Template

func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	models.Templates, models.DummyError = template.ParseGlob("templates/*")
	utils.HandleErr(models.DummyError, "Error parsing glob from /templates", "Parsed templates fom /templates")

	err02 := godotenv.Load(".env")
	utils.HandleErr(err02, "Failed to export env variables from .env file", "Exported env variables from .env file")

	ctx := context.Background()
	models.DummyError = nil
	models.PubSubClient, models.DummyError = pubsub.NewClient(ctx, "manavsethi")
	utils.HandleErr(models.DummyError, "Error creating Pubsub client", "Created Pubsub client!")

}

func main() {

	err := StartServer()
	utils.HandleErr(err, "Error starting HTTP server at port: "+os.Getenv("PORT"), "HTTP server listening at port: "+os.Getenv("PORT"))

}

func StartServer() error {

	FileServer := StartFileServer()
	Multiplexer := http.NewServeMux()

	Multiplexer.HandleFunc("/", controllers.Home)
	Multiplexer.H
	Multiplexer.Handle("/chat", websocket.Handler(controllers.Chat))

	Multiplexer.HandleFunc("/resume", controllers.Resume)

	Multiplexer.Handle("/storage/", http.StripPrefix("/storage/", FileServer))

	err := http.ListenAndServe(":"+os.Getenv("PORT"), Multiplexer)
	if err != nil {
		return err
	}
	return nil
}

func StartFileServer() http.Handler {
	FileServer := http.FileServer(http.Dir("/storage"))
	return FileServer
}
