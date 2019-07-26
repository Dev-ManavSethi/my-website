package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Dev-ManavSethi/my-website/models"

	"github.com/Dev-ManavSethi/my-website/utils"
)

func ChatPage(ResponseWriter http.ResponseWriter, Request *http.Request) {

}

func About(ResponseWriter http.ResponseWriter, Request *http.Request) {
	fmt.Fprintln(ResponseWriter, "Coming soon!")
}

func Projects(ResponseWriter http.ResponseWriter, Request *http.Request) {

	resp, err := http.Get(os.Getenv("GITHUB_REPOS_API_URL"))
	if err != nil {
		fmt.Fprintln(ResponseWriter, "Error: ", err, "Refresh or try later!")
	} else {

		var ReposFromGitHub []models.GitRepo

		ResponseBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(ResponseWriter, "Error: ", err, "Refresh or try later!")
		} else {
			err := json.Unmarshal(ResponseBytes, &ReposFromGitHub)
			if err != nil {
				fmt.Fprintln(ResponseWriter, "Error: ", err, "Refresh or try later!")
			} else {
				err := models.Templates.ExecuteTemplate(ResponseWriter, "projects.html", ReposFromGitHub)
				if err != nil {
					fmt.Fprintln(ResponseWriter, "Error: ", err, "Refresh or try later!")
				}
			}
		}

		defer resp.Body.Close()
	}

}
func Home(ResponseWriter http.ResponseWriter, Request *http.Request) {

	IPaddress := utils.GetUserIP(Request)

	models.GlobalMutex.Lock()
	User := models.Chats[IPaddress]
	User.VisitCount++

	if User.VisitCount > 1 {
		User.VisitMoreThanOnce = true
	}

	models.Chats[IPaddress] = User
	models.GlobalMutex.Unlock()

	err := models.Templates.ExecuteTemplate(ResponseWriter, "home.html", models.Chats[IPaddress])
	utils.HandleErr(err, "Unable to execute template home.html", "")

}

func ResumeUpload(ResponseWriter http.ResponseWriter, Request *http.Request) {

	if Request.Method == http.MethodPost {

		FileReader, err := os.OpenFile("storage/pdf/resume.pdf", os.O_CREATE|os.O_RDWR, 0655)
		if err != nil {
			log.Println(err)
			ResponseWriter.WriteHeader(http.StatusInternalServerError)
		} else {
			err := Request.ParseForm()
			if err != nil {
				log.Println(err)
				ResponseWriter.WriteHeader(http.StatusInternalServerError)
			} else {
				file, _, err := Request.FormFile("resume")
				if err != nil {
					log.Println(err)
					ResponseWriter.WriteHeader(http.StatusInternalServerError)
				} else {
					_, err := io.Copy(FileReader, file)
					if err != nil {
						log.Println(err)
						ResponseWriter.WriteHeader(http.StatusInternalServerError)
					} else {
						ResponseWriter.WriteHeader(http.StatusOK)
						fmt.Fprint(ResponseWriter, "Done!")
					}
				}
			}
		}
	}

	if Request.Method == http.MethodGet {
		err := models.Templates.ExecuteTemplate(ResponseWriter, "upload_resume.html", nil)
		utils.HandleErr(err, "Unable to execute template upload_resume.html", "")
	}

}

func Resume(ResponseWriter http.ResponseWriter, Request *http.Request) {

	// HTTPresponse, error01 := http.Get(os.Getenv("RESUME_URL"))
	// utils.HandleErr(error01, "Error getting resume from: "+os.Getenv("RESUME_URL"), "")

	// ResumeFileWriter, error02 := os.Create("./storage/pdf/resume.pdf")
	// utils.HandleErr(error02, "Error creating / getting /storage/resume.pdf", "")

	// _, error03 := io.Copy(ResumeFileWriter, HTTPresponse.Body)
	// defer HTTPresponse.Body.Close()
	// // defer ResumeFileWriter.Close()

	// utils.HandleErr(error03, "Error copying from resonse body to file", "Updated resume!")

	http.ServeFile(ResponseWriter, Request, "storage/pdf/resume.pdf")

}
