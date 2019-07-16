package controllers

import (
	"github.com/Dev-ManavSethi/my-website/models"
	"github.com/Dev-ManavSethi/my-website/utils"
	"log"
	"net/http"
	"time"
)



func Chat(w http.ResponseWriter, r *http.Request){


	if r.Method==http.MethodGet {

		IPAddress := utils.GetUserIP(r)
		log.Println(IPAddress)


		UserExists := utils.CheckChatUserExists(IPAddress)

		if UserExists{
			err:= models.Templates.ExecuteTemplate(w, "chat.html", models.Chats[IPAddress])
			utils.HandleErr(err, "Error executing chat.html", "")

		} else{

			name := r.URL.Query().Get("name")
			if name==""{

				err := models.Templates.ExecuteTemplate(w, "chat.html", nil)
				utils.HandleErr(err, "Error executing chat.html", "")

			} else {
				utils.RegisterChatUser(IPAddress, name)
				err := models.Templates.ExecuteTemplate(w, "chat.html", models.Chats[IPAddress])
				utils.HandleErr(err, "Error executing chat.html", "")

				error := utils.BackupChats()
				if error!=nil{
					log.Println("Unable to nackup chats")
					log.Fatalln(error)
				}

			}





		}


	}



	if r.Method == http.MethodPost{

		err := r.ParseForm()
		if err!=nil{
			//handle error
		}

		name := r.FormValue("name")
		message := r.FormValue("message")
		time := time.Now().Unix()

		IncomingMessage := models.ChatMessage{
			Name:name,
			Message:message,
			Time:time,
		}

		IPAddress := utils.GetUserIP(r)
		log.Println(IPAddress)

		UserExists := utils.CheckChatUserExists(IPAddress)

		if UserExists{

		User := models.Chats[IPAddress]

		User.Chats = append(User.Chats, IncomingMessage)

		models.Chats[IPAddress]= User
			error := utils.BackupChats()
			if error!=nil{
				log.Println("Unable to nackup chats")
				log.Fatalln(error)
			}


		} else {

			utils.RegisterChatUser(IPAddress, name)
			User := models.Chats[IPAddress]

			User.Chats = append(User.Chats, IncomingMessage)

			models.Chats[IPAddress]= User

			error := utils.BackupChats()
			if error!=nil{

				log.Println("Unable to nackup chats")
			log.Fatalln(error)
			}






		}
	}



}
