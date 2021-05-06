package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
	"testing"
	"thrift-go-demo/gen-go/sample"
)

func TestSayHello(t *testing.T) {
	trans, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}

	defer trans.Close()

	protocolFactory := thrift.NewTCompactProtocolFactory()

	client := sample.NewGreeterClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", "localhost", ":", "9090", " ", err)
		os.Exit(1)
	}

	value0 := &sample.User{
		Id:      1,
		Name:    "sunchp",
		Avatar:  "1231",
		Address: "北京",
		Mobile:  "12321",
	}

	fmt.Print(client.SayHello(value0))
	fmt.Print("\n")
}
