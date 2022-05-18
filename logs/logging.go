
package logging

import (
    "log"
    //"log/syslog"
    "os"
)

//Function to handle error logging to file globally to file errorlog
func Logerror(err error, msg string) {
    // log to custom file
    LOG_FILE := "./logs/errorlog"
    // open log file
    logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        log.Panic(err)
    }
    defer logFile.Close()

    // Set log output file)
    log.SetOutput(logFile)

    //log date-time, filename, and line number
    log.SetFlags(log.Lshortfile | log.LstdFlags)


	log.Println(msg, err)
}

//Function to handle the logging of information globally to file infolog
func Loginfo(msg string){
	LOG_FILE := "./logs/infolog"
    // open log file
    logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        log.Panic(err)
    }
    defer logFile.Close()

    // Set log out put and enjoy :)
    log.SetOutput(logFile)

    // optional: log date-time, filename, and line number
    log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println(msg)
}


//Må sjekke om denne her brukes
func Logerrorinfo(msg string) {
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

    //log date-time, filename, and line number
    log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println(msg)
}