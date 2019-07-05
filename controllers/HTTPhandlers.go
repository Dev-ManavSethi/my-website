package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/Dev-ManavSethi/my-website/models"

	"github.com/Dev-ManavSethi/my-website/utils"
)

func ChatPage(ResponseWriter http.ResponseWriter, Request *http.Request) {

}

func About(ResponseWriter http.ResponseWriter, Request *http.Request) {

}

func Projects(ResponseWriter http.ResponseWriter, Request *http.Request) {

	resp, err := http.Get(os.Getenv("GITHUB_REPOS_API_URL"))

}
func Home(ResponseWriter http.ResponseWriter, Request *http.Request) {

	err := models.Templates.ExecuteTemplate(ResponseWriter, "home.html", nil)
	utils.HandleErr(err, "Unable to execute template", "")

}

func Resume(ResponseWriter http.ResponseWriter, Request *http.Request) {

	HTTPresponse, error01 := http.Get(os.Getenv("RESUME_URL"))
	utils.HandleErr(error01, "Error getting resume from: "+os.Getenv("RESUME_URL"), "")

	ResumeFileWriter, error02 := os.Create("./storage/pdf/resume.pdf")
	utils.HandleErr(error02, "Error creating / getting /storage/resume.pdf", "")

	_, error03 := io.Copy(ResumeFileWriter, HTTPresponse.Body)
	defer HTTPresponse.Body.Close()
	defer ResumeFileWriter.Close()

	utils.HandleErr(error03, "Error copying from resonse body to file", "Updated resume!")

	http.ServeFile(ResponseWriter, Request, "storage/pdf/resume.pdf")

}
