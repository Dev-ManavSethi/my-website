package controllers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

//mongodb://manav:manavsethi1@ds251507.mlab.com:51507/heroku_qfjcjl3j

func ConnectToMongo() (*mongo.Client, error){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + os.Getenv("MONGO_USER") + os.Getenv("MONGO_PASS") + "@" +os.Getenv("MONGO_HOST") + "/"+ os.Getenv("MONGO_DB_NAME") + "?retryWrites=true&w=majority"))
	if err!=nil{
		return  nil, err
	}

	err1 := client.Connect(context.TODO())
	if err1!=nil{
		return nil, err1
	}
	err2 := client.Ping(context.TODO(), nil)
	if err!=nil{
		return  nil, err2
	}

	return client, nil


}
