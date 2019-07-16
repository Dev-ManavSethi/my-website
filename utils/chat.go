package utils

import (
	"bytes"
	"encoding/gob"
	"github.com/Dev-ManavSethi/my-website/models"
	"os"
	"time"
)


func LoadChatsFromFile(filename string) (map[string]models.ChatUser,error){

	file, err := os.OpenFile("chats.file", os.O_RDWR|os.O_CREATE, 0755)
	if err!=nil{
		return nil, err
	}

	maps := make(map[string]models.ChatUser)


decoder := gob.NewDecoder(file)
err2 := decoder.Decode(&maps)
if err2!=nil{
	return nil, err
}

//log.Println(maps)
 return maps ,nil



}
func BackupChats()error{


	file, err := os.OpenFile("chats.file", os.O_RDWR|os.O_CREATE, 0755)
	if err!=nil{
return err
	}

	defer  file.Close()



var ByteChats bytes.Buffer
	encoder := gob.NewEncoder(&ByteChats)
	err2 := encoder.Encode(models.Chats)
	if err2!=nil{
		return err
	}

	_, er := file.Write(ByteChats.Bytes())
	if er!=nil{
		return er
	}


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

	models.Chats[IPAddress] = models.ChatUser{
		Name:name,
		IP:IPAddress,
		Chats:Chats,
	}

}
