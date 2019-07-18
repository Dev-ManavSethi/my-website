package utils

import (
	"bytes"
	"encoding/gob"
	"github.com/Dev-ManavSethi/my-website/models"
	"log"
	"os"
	"time"
)


func LoadChatsFromFile(filename string) (map[string]models.User,error){

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err!=nil{
		return nil, err
	}

	maps := make(map[string]models.User)


	decoder := gob.NewDecoder(file)
	models.GlobalMutex.Lock()
	err2 := decoder.Decode(&maps)
	models.GlobalMutex.Unlock()
	if err2!=nil{
	return nil, err
	}


//log.Println(maps)
 return maps ,nil



}
func BackupChats()error{


	file, err := os.OpenFile(os.Getenv("CHATS_DB"), os.O_RDWR|os.O_CREATE, 0755)
	if err!=nil{
	return err
	}

	defer  file.Close()

	log.Println("Starting Chat backup")

	var ByteChats bytes.Buffer

	encoder := gob.NewEncoder(&ByteChats)

	models.GlobalMutex.Lock()
	err2 := encoder.Encode(models.Chats)
	models.GlobalMutex.Unlock()
	if err2!=nil{
		return err
	}

	models.GlobalMutex.Lock()
	_, er := file.Write(ByteChats.Bytes())
	models.GlobalMutex.Unlock()
	if er!=nil{
		return er
	}

	log.Println("Chat backup done!")


return  nil
}

func CheckChatUserExists(IPAddress string) bool{
	var UserExists bool = false

	for ip, _ := range models.Chats{

		if ip==IPAddress{
			UserExists=true
		}
	}

	return UserExists
}

func RegisterChatUser(IPAddress, name string) {
	var Chats []models.ChatMessage

	Chats = append(Chats, models.ChatMessage{
		Name:"Manav",
		Message:"Hi " + name +", How may I help you?",
		Time:time.Now().Unix(),
	})


	models.GlobalMutex.Lock()
	models.Chats[IPAddress] = models.User{
		Name:name,
		IP:IPAddress,
		Chats:Chats,
	}
	models.GlobalMutex.Unlock()

}
