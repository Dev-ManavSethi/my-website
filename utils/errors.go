package utils

import "log"

func HandleErr(err error, ErrorMessage, SuccessMesaage string) {

	if err != nil {
		log.Println(ErrorMessage)
		log.Fatalln(err)
	} else if SuccessMesaage != "" {
		log.Println(SuccessMesaage)
	}

}
