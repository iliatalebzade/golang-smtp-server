package server

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type SMTPServer struct {
	l *log.Logger
}

func NewSmtpServer(logger *log.Logger) *SMTPServer {
	return &SMTPServer{l: logger}
}

func receiveEmail(client net.Conn, logger *log.Logger) {
	emailHeaders := make(map[string]string)
	var emailBody strings.Builder

	// Read the email headers from the client until a blank line is received
	reader := bufio.NewReader(client)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}
		line = strings.TrimSpace(line)

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			emailHeaders[key] = value
		}
	}

	// Read the email body from the client until a line with a single period (.) is received
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == ".\r\n" {
			break
		}
		emailBody.WriteString(strings.TrimSpace(line) + "\n")
	}

	// Print the received email headers and body
	logger.Println("Received email headers:")
	for key, value := range emailHeaders {
		logger.Printf("%s: %s\n", key, value)
	}
	logger.Println("Received email body:")
	logger.Println(emailBody.String())
}

func (s *SMTPServer) StartSMTPServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:25")
	if err != nil {
		s.l.Println("Failed to start SMTP server:", err)
		return
	}
	s.l.Println("SMTP server started on localhost:25")

	for {
		client, err := listener.Accept()
		if err != nil {
			s.l.Println("Failed to accept incoming connection:", err)
			continue
		}
		s.l.Println("Incoming connection from:", client.RemoteAddr().String())

		client.Write([]byte("220 localhost Simple SMTP Server\r\n"))

		scanner := bufio.NewScanner(client)
		for scanner.Scan() {
			line := scanner.Text()
			command := strings.Split(line, " ")[0]
			s.l.Println("--------------------------------------------------------------------")
			s.l.Println(command)
			s.l.Println("--------------------------------------------------------------------")
			switch strings.ToUpper(command) {
			case "HELO", "EHLO":
				client.Write([]byte("250 Hello " + strings.Split(line, " ")[1] + "\r\n"))
			case "MAIL":
				client.Write([]byte("250 OK\r\n"))
			case "RCPT":
				client.Write([]byte("250 OK\r\n"))
			case "DATA":
				client.Write([]byte("354 Start mail input; end with <CRLF>.<CRLF>\r\n"))
				receiveEmail(client, s.l)
				client.Write([]byte("250 OK\r\n"))
			case "QUIT":
				client.Write([]byte("221 Bye\r\n"))
				client.Close()
				break
			default:
				client.Write([]byte("500 Command not recognized\r\n"))
			}
		}

		client.Close()
	}
}
