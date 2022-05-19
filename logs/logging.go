package logging

import (
	"log"
	//"log/syslog"
	"os"
)

//Function to handle logging of errors to errorlog file with message
func Logerror(err error, msg string) {
	// log to custom file
	LOG_FILE := "./logs/errorlog"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log output file
	log.SetOutput(logFile)

	// log date-time, filename
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println(msg, err)
}

//Function to handle information logging to infofile
func Loginfo(msg string) {
	LOG_FILE := "./logs/infolog"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log output
	log.SetOutput(logFile)

	//log date-time, filename
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println(msg)
}

//Function to handle error message display to file. 
func Logerrorinfo(msg string) {
	// log to custom file
	LOG_FILE := "./logs/errorlog"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log output file :)
	log.SetOutput(logFile)

	//log date-time, filename
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println(msg)
}