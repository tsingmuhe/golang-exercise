package main

import (
	"encoding/json"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"thrift-go-demo/gen-go/sample"
)

type GreeterServiceImpl struct {
}

func (p *GreeterServiceImpl) SayHello(user *sample.User) (r *sample.Response, err error) {
	strJson, _ := json.Marshal(user)
	fmt.Printf("message from client: %s\n", strJson)
	return &sample.Response{ErrCode: 0, ErrMsg: "success", Data: map[string]string{"User": string(strJson)}}, nil
}

func (p *GreeterServiceImpl) GetUser(uid int32) (r *sample.Response, err error) {
	fmt.Printf("message from client: %v\n", uid)
	return &sample.Response{ErrCode: 1, ErrMsg: "user not exist."}, nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		panic(err)
	}

	processor := sample.NewGreeterProcessor(&GreeterServiceImpl{})
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTCompactProtocolFactory(),
	)

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
