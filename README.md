# Tcp-chat

A simple TCP chat server in go

### Build

`go get ./...
go run main.go`

### Build Docker image

`docker build . -t tcp-chat`

#### Run Docker container

`docker run -p 8888:8888 tcp-chat`

### Usage
Run the server either using docker conatiner or thorugh source code

Use telnet to connect to server 

`telnet localhost 8888`

### Commands

```/nick nickname := sets nickname of the user

/join grp-name := join the grp-name if exists else create a new group

/msg message := sends message to all users of current grp

/rooms := list all the rooms availaible

/quit := quit
```
