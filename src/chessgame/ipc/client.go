package ipc

import "encoding/json"

type IpcClient struct {
	connect chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (response *Response, err error) {
	request := &Request{Method: method, Params: params}
	var b []byte
	b, err = json.Marshal(request)
	if err != nil {
		return
	}
	client.connect <- string(b)
	str := <-client.connect
	var resp Response
	err = json.Unmarshal([]byte(str), resp)
	response = &resp
	return
}

func (client *IpcClient) Close() {
	client.connect <- "CLOSE"
}
