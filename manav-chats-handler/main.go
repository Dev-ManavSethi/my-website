package main

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

type ChatMessage struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Address string `json:"address"`
}

var (
	Pubsubclient *pubsub.Client
	err          error
)

func HandleErr(err error, ErrorMessage, SuccessMesaage string) {

	if err != nil {
		log.Println(ErrorMessage)
		log.Fatalln(err)
	} else if SuccessMesaage != "" {
		log.Println(SuccessMesaage)
	}

}

func init() {
	err02 := godotenv.Load(".env")
	HandleErr(err02, "ERR LOADING ENV VARS", "LOADED ENV VARS")

	Pubsubclient, err = pubsub.NewClient(context.Background(), "manavsethi")
	HandleErr(err, "Error creating pub sub client", "Created pub sub client!")
}

func main() {

	err := Pubsubclient.Subscription("manav_sub").Receive(context.TODO(), func(ctx context.Context, msg *pubsub.Message) {

		var IncomingMessage ChatMessage
		err2 := json.Unmarshal(msg.Data, &IncomingMessage)
		if err2 != nil {

			log.Fatalln(err)

		} else {
			RemoteAddress := IncomingMessage.Address

			var ReplyFromManav ChatMessage
			ReplyFromManav.Name = "Manav"
			ReplyFromManav.Message = "Hi! This is automated reply"
			ReplyFromManav.Address = ""

			JSONbytes, err := json.Marshal(ReplyFromManav)
			if err != nil {
				log.Fatalln(err)
			} else {

				exists, err := Pubsubclient.Topic(RemoteAddress).Exists(context.TODO())
				if err != nil {
					log.Fatalln(err)
				}
				if exists {

					result := Pubsubclient.Topic(RemoteAddress).Publish(context.TODO(), &pubsub.Message{
						Data: JSONbytes,
					})

					_, err := result.Get(context.TODO())
					if err != nil {

					} else {
						log.Println("Replied with automated message!")
					}
				}

				if !exists {
					Topic, err := Pubsubclient.CreateTopic(context.TODO(), RemoteAddress)
					if err != nil {
						log.Fatalln(err)
					} else {
						result := Topic.Publish(context.TODO(), &pubsub.Message{
							Data: JSONbytes,
						})
						_, err := result.Get(context.TODO())
						if err != nil {

						} else {
							log.Println("Replied with automated message!")
						}
					}

				}

			}

		}
	})
	if err != nil {
		log.Println(err)
	}

}
