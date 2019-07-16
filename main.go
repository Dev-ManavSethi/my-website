package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Dev-ManavSethi/my-website/models"

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

	models.DummyError= nil
	models.MongoDBclient, models.DummyError = controllers.ConnectToMongo()
	utils.HandleErr(models.DummyError, "Mongo err", "Connected to MongoDB")

}

func main() {

	err := StartServer()
	utils.HandleErr(err, "Error starting HTTP server at port: "+os.Getenv("PORT"), "HTTP server listening at port: "+os.Getenv("PORT"))

}

func StartServer() error {

	FileServer := http.FileServer(http.Dir("storage"))
	Multiplexer := http.NewServeMux()

	Multiplexer.HandleFunc("/", controllers.Home) //done
	Multiplexer.HandleFunc("/about", controllers.About)
	Multiplexer.HandleFunc("/resume", controllers.Resume)     //done
	Multiplexer.HandleFunc("/projects", controllers.Projects) //done

	Multiplexer.HandleFunc("/chat", controllers.Chat)


	Multiplexer.Handle("/storage/", http.StripPrefix("/storage/", FileServer))

	log.Println("Listening at : " + os.Getenv("PORT"))

	err := http.ListenAndServe(":"+os.Getenv("PORT"), Multiplexer)
	if err != nil {
		return err
	}
	return nil
}

// func StartFileServer() http.Handler {

// }
