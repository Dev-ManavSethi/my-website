package controllers

import (
	"context"
	"github.com/Dev-ManavSethi/my-website/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strings"
	"time"
)

//func Chat(WebSocketConn *websocket.Conn) {
//
//	log.Println("Chat initiated!")
//
//	websocket.JSON.Send(WebSocketConn, models.ChatMessage{
//		Name:    "Manav",
//		Message: "Hi! How can I help?",
//		Address: "",
//	})
//
//	go Reply(WebSocketConn)
//
//	ListenMessages(WebSocketConn)
//
//	for {
//
//	}
//
//	log.Println("Closing WS connection")
//	defer WebSocketConn.Close()
//
//}
//
//func ListenMessages(WebSocketConn *websocket.Conn) {
//
//	for {
//		var IncomingMessage *models.ChatMessage
//		err := websocket.JSON.Receive(WebSocketConn, &IncomingMessage)
//
//		if err != nil {
//			log.Fatalln(err)
//
//		} else {
//			log.Println("Recieved chat message!")
//			IncomingMessage.Address = "to-be-replaced-by-client-addr"
//
//			JSONbytes, err := json.Marshal(*IncomingMessage)
//			if err != nil {
//
//				log.Fatalln(err)
//
//			} else {
//
//				result := models.PubSubClient.Topic("manav").Publish(context.TODO(), &pubsub.Message{
//					Data: JSONbytes,
//				})
//				_, err := result.Get(context.TODO())
//				if err != nil {
//					log.Fatalln(err)
//				} else {
//					//send confirmation of reciept
//					log.Println("Sent message to manav")
//				}
//			}
//		}
//	}
//
//}
//
//func Reply(WebSocketConn *websocket.Conn) {
//
//	for {
//		//Topic := models.PubSubClient.Topic("to-be-replaced-by-client-addr")
//		sub := models.PubSubClient.Subscription("")
//		cctx, _ := context.WithCancel(context.TODO())
//		err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
//			msg.Ack()
//			fmt.Printf("Recieved reply from Manav!")
//
//			err2 := websocket.JSON.Send(WebSocketConn, string(msg.Data))
//
//			if err2 != nil {
//				//SendErrorMessage(WebSocketConn, err2)
//
//			}
//		})
//
//		if err != nil {
//
//			//	SendErrorMessage(WebSocketConn, err)
//		}
//
//	}
//}


func Chat(w http.ResponseWriter, r *http.Request){




	if r.Method == http.MethodPost{


		err := r.ParseForm()
		if err!=nil{

		}
		name := r.FormValue("name")
		message := r.FormValue("message")


		IPAddress := r.Header.Get("X-Real-Ip")
		if IPAddress == "" {
			IPAddress = r.Header.Get("X-Forwarded-For")
		}
		if IPAddress == "" {
			IPAddress = r.RemoteAddr
		}

		Addresses := strings.Split(IPAddress, ":")
		IPAddress = Addresses[0]

		var Message models.ChatMessage
		Message.Name = name
		Message.Message = message
		Message.Time = time.Now().Unix()

		//check if user exists or not

		result:=  models.MongoDBclient.Database("heroku_qfjcjl3j").Collection("chats").FindOne(context.TODO(), bson.M{
			"ip": IPAddress})

		if result.Err()!=nil{
			//user dne
			//add user
			User := models.ChatUser{}
			User.IP = IPAddress
			User.Name = name
			User.Chats = append(User.Chats, models.ChatMessage{
				Name:name,
				Message:message,
				Time:time.Now().Unix()})

		_, err := 	models.MongoDBclient.Database("heroku_qfjcjl3j").Collection("chats").InsertOne(context.TODO(), User)
		if err!=nil{
			log.Fatal(err)
		}




		} else {

			//user exists in db
		_ ,err :=	models.MongoDBclient.Database("heroku_qfjcjl3j").Collection("chats").UpdateOne(context.TODO(), bson.M{"ip": IPAddress}, bson.M{"$push": bson.M{"chats": Message}})
		if err!=nil{
			log.Fatal(err)
		}
			_ ,err2 :=	models.MongoDBclient.Database("heroku_qfjcjl3j").Collection("chats").UpdateOne(context.TODO(), bson.M{"ip": IPAddress}, bson.M{"$set": bson.M{"name": name}})
			if err2!=nil{
				log.Fatal(err2)
			}





		}

		//get updated user

		//result2:=  models.MongoDBclient.Database("heroku_qfjcjl3j").Collection("chats").FindOne(context.TODO(), bson.M{
		//	"ip": IPAddress})
		//
		//if result2.Err()!=nil{
		//	//handle error
		//}
		//
		//var User models.ChatUser
		//
		//err2 := result.Decode(&User)
		//if err2!=nil{
		//	//handle error
		//}
		//err3 := models.Templates.ExecuteTemplate(w, "chat.html", User)
		//if err3!=nil{
		//
		//}




	}


	if r.Method==http.MethodGet {

		IPAddress := r.Header.Get("X-Real-Ip")
		if IPAddress == "" {
			IPAddress = r.Header.Get("X-Forwarded-For")
		}
		if IPAddress == "" {
			IPAddress = r.RemoteAddr
		}


		Addresses := strings.Split(IPAddress, ":")
		IPAddress = Addresses[0]

		log.Println(IPAddress)

		//find by ip if exists or not
		result := models.MongoDBclient.Database("heroku_qfjcjl3j").Collection("chats").FindOne(context.TODO(), bson.M{
			"ip": IPAddress})

		if result.Err() != nil {
			//handle error
			//add user

			log.Println("Adding new Chat User: ip: " + IPAddress)
			//add user
			User := models.ChatUser{}
			User.IP = IPAddress
			User.Name = "Guest"
			User.Chats = append(User.Chats, models.ChatMessage{
				Name:"Manav",
				Message:"Hi! How may I help you?",
				Time:time.Now().Unix()})



			_, err := 	models.MongoDBclient.Database("heroku_qfjcjl3j").Collection("chats").InsertOne(context.TODO(), User)
			if err!=nil{
				log.Fatal(err)
			}

			log.Println("New chat user added!")

			//get user

			//
			//result := models.MongoDBclient.Database("heroku_qfjcjl3j").Collection("chats").FindOne(context.TODO(), bson.M{
			//	"ip":IPAddress})
			//
			//
			//var User2 models.ChatUser
			//
			//err2 := result.Decode(&User2)
			//if err2 != nil {
			//	//handle error
			//}
			err4 := models.Templates.ExecuteTemplate(w, "chat.html", User)
			if err4 != nil {

			}

		} else {

			//user already exists


			var User models.ChatUser

			err := result.Decode(&User)
			if err != nil {
				//handle error
			}
			log.Println(User.Name + " is back on chat!")
			err2 := models.Templates.ExecuteTemplate(w, "chat.html", User)
			if err2 != nil {

			}
		}

	}

}
