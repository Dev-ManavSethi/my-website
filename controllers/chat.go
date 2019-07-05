package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/Dev-ManavSethi/my-website/models"
	"golang.org/x/net/websocket"
)

func Chat(WebSocketConn *websocket.Conn) {

	log.Println("Chat initiated!")

	websocket.JSON.Send(WebSocketConn, models.ChatMessage{
		Name:    "Manav",
		Message: "Hi! How can I help?",
		Address: "",
	})

	go Reply(WebSocketConn)

	ListenMessages(WebSocketConn)

	for {

	}

	log.Println("Closing WS connection")
	defer WebSocketConn.Close()

}

func ListenMessages(WebSocketConn *websocket.Conn) {

	for {
		var IncomingMessage *models.ChatMessage
		err := websocket.JSON.Receive(WebSocketConn, &IncomingMessage)

		if err != nil {
			log.Fatalln(err)

		} else {
			log.Println("Recieved chat message!")
			IncomingMessage.Address = "to-be-replaced-by-client-addr"

			JSONbytes, err := json.Marshal(*IncomingMessage)
			if err != nil {

				log.Fatalln(err)

			} else {

				result := models.PubSubClient.Topic("manav").Publish(context.TODO(), &pubsub.Message{
					Data: JSONbytes,
				})
				_, err := result.Get(context.TODO())
				if err != nil {
					log.Fatalln(err)
				} else {
					//send confirmation of reciept
					log.Println("Sent message to manav")
				}
			}
		}
	}

}

func Reply(WebSocketConn *websocket.Conn) {

	for {
		Topic := models.PubSubClient.Topic("to-be-replaced-by-client-addr").
		sub := models.PubSubClient.Subscription("")
		cctx, _ := context.WithCancel(context.TODO())
		err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
			msg.Ack()
			fmt.Printf("Recieved reply from Manav!")

			err2 := websocket.JSON.Send(WebSocketConn, string(msg.Data))

			if err2 != nil {
				//SendErrorMessage(WebSocketConn, err2)

			}
		})

		if err != nil {

			//	SendErrorMessage(WebSocketConn, err)
		}

	}
}
