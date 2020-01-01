package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string
	Params string
}

type Response struct {
	Code string
	Body string
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			request := <-c
			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}
			response := server.Handle(req.Method, req.Params)
			result, err := json.Marshal(response)
			c <- string(result)
		}
		fmt.Println("Session closed.")
	}(session)

	fmt.Println("A new session has been created successfully.")
	return session
}
