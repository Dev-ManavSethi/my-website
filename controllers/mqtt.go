package controllers

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConnectCloudMQTT(ClientID string) (mqtt.Client, error) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", os.Getenv("MQTT_HOST"), os.Getenv("MQTT_PORT")))
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))

	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
	opts.SetClientID(ClientID)

	client := mqtt.NewClient(opts)

	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		return nil, err
	}

	log.Println("Connect to Cloud MQTT, ClientID: " + ClientID)
	return client, nil
}
