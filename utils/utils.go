package utils

import "log"

func FatalIfErrorNotNil(err error, message string, Return bool) {

	if err != nil {
		log.Println(message)
		log.Fatalln(err)

	}

	if Return {
		return
	} else {

	}
}
