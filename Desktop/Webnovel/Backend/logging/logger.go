package logging

import (
	"fmt"
	"log"
	"os"
)

func InitLogger() {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Cant init log file")
	}

	log.SetOutput(logFile)
}
