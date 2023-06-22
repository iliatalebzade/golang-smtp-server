package main

import (
	"log"
	"smpt_server/logger"
	"smpt_server/server"
)

func main() {
	// Setup the logger
	logger, err := logger.SetupLogger()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// setup a new server instance
	s := server.NewSmtpServer(logger)

	// start the SMTP server
	s.StartSMTPServer()
}
