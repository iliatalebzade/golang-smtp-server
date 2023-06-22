# golang-smtp-server

A simple SMTP server written in Go that catches and logs mails sent to the port 25 of your host.

## About

This SMTP server is designed to receive and log emails sent to the specified port (default: 25) on your host. It provides a basic implementation of an SMTP server and captures the email headers and body for further processing or logging.

## How to Run

To run the SMTP server, follow the steps below:

1. Clone the project:
git clone <repository-url>

2. Change into the directory of the cloned source code:
cd golang-smtp-server

3. Build the Docker image:
sudo docker build -t smtp-server .

4. Run the Docker container, mapping port 25 to the host:
sudo docker run -p 25:25 smtp-server


> Note: Make sure you have Docker installed and running on your host machine.

## Viewing Logs

To view the logs generated by the SMTP server, you can follow these steps:

1. Attach to the shell of your running container:
