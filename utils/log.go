package utils

import (
	"log"
	"os"
)

func LogToFile(filename string) error{

	log.Println("Setting log file to "+filename)
	f, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)
	if err != nil {
		return err
	}
//	defer f.Close()

	log.SetOutput(f)

	return nil
}
