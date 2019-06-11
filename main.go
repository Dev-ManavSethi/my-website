package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"./utils"
)

var tpl *template.Template

func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var error01 error

	tpl, error01 = template.ParseGlob("./templates/*")

	utils.FatalIfErrorNotNil(error01, "Failed to Parse glob templates from ./templates", true)

	error02 := godotenv.Load()
	utils.FatalIfErrorNotNil(error02, "Error loading env variables", true)

}

func main() {

	FileServer := http.FileServer(http.Dir("/storage"))

	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/resume", resume)

	multiplexer.Handle("/storage/", http.StripPrefix("/storage/", FileServer))

	error01 := http.ListenAndServe(":"+os.Getenv("PORT"), multiplexer)
	utils.FatalIfErrorNotNil(error01, "Error staring server at port "+os.Getenv("PORT"), true)
}

func resume(ResponseWriter http.ResponseWriter, Request *http.Request) {

	HTTPresponse, error01 := http.Get(os.Getenv("RESUME_URL"))
	utils.FatalIfErrorNotNil(error01, "Error getting response from resume url", false)

	ResumeFileWriter, error02 := os.Create("./storage/pdf/resume.pdf")
	utils.FatalIfErrorNotNil(error02, "Error creating resume pdf file", false)

	_, error03 := io.Copy(ResumeFileWriter, HTTPresponse.Body)
	defer HTTPresponse.Body.Close()
	defer ResumeFileWriter.Close()

	utils.FatalIfErrorNotNil(error03, "Error copying from resonse body to file", false)

	log.Println("Resume updated")

	http.ServeFile(ResponseWriter, Request, "storage/pdf/resume.pdf")
	log.Println("Resume shown")

}
