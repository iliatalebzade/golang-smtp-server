# Use the official GoLang base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Install air for hot-reloading
RUN apk update && apk add --no-cache git
RUN go mod init smpt_server
# RUN go get -u github.com/cosmtrek/air

# Copy the GoLang application source code to the container
COPY . .

# Expose a port (if your application listens on a specific port)
EXPOSE 25

# Specify the command to run the "air" tool when the container starts
CMD ["go", "run", "main.go"]
