package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"golang.org/x/net/websocket"

	"github.com/Dev-ManavSethi/my-website/models"

	"github.com/Dev-ManavSethi/my-website/controllers"
	"github.com/Dev-ManavSethi/my-website/utils"
)

func init() {

	err02 := godotenv.Load(".env")
	utils.HandleErr(err02, "Failed to export env variables from .env file", "Exported env variables from .env file")

	//
	// err := utils.LogToFile(os.Getenv("LOG_FILE"))
	// utils.HandleErr(err, "Error setting log file", "Log file set!")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	models.Templates, models.DummyError = template.ParseGlob("templates/*")
	utils.HandleErr(models.DummyError, "Error parsing glob from /templates", "Parsed templates fom /templates")

	models.Chats = make((map[string]models.User))

	models.DummyError = nil
	// models.Chats, models.DummyError = utils.LoadChatsFromFile(os.Getenv("CHATS_DB"))
	// utils.HandleErr(models.DummyError, "Error loading chats from db/chats.db", "Loaded chats from db/chats.db")
	// //log.Println(models.Chats)

}

func main() {

	go func() {

		var GraceSignal = make(chan os.Signal)

		signal.Notify(GraceSignal, syscall.SIGTERM)
		signal.Notify(GraceSignal, syscall.SIGINT)

		sig := <-GraceSignal
		log.Println("Signal recieved", sig)
		utils.BackupChats()
		log.Println("Exiting app")
		log.Println()
		os.Exit(0)

	}()

	err := StartServer()
	utils.HandleErr(err, "Error starting HTTP server at port: "+os.Getenv("PORT"), "HTTP server listening at port: "+os.Getenv("PORT"))

}

func StartServer() error {

	FileServer := http.FileServer(http.Dir(os.Getenv("STORAGE_DIR")))

	Multiplexer := http.NewServeMux()

	Multiplexer.HandleFunc("/", controllers.Home) //done
	Multiplexer.HandleFunc("/about", controllers.About)
	Multiplexer.HandleFunc("/resume", controllers.Resume)     //done, update resume + env resume link
	Multiplexer.HandleFunc("/resume/upload", controllers.ResumeUpload)
	Multiplexer.HandleFunc("/projects", controllers.Projects) //done
	Multiplexer.HandleFunc("/chat", controllers.Chat)
	Multiplexer.Handle("/chatws", websocket.Handler(controllers.ChatWS))

	Multiplexer.Handle("/storage/", http.StripPrefix("/storage/", FileServer))

	log.Println("Listening at : " + os.Getenv("PORT"))

	models.HTTPserver = &http.Server{
		Handler: Multiplexer,
	}

	l, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		return err
	}

	err = models.HTTPserver.Serve(l)
	if err != nil {
		return err
	}

	return nil
}
