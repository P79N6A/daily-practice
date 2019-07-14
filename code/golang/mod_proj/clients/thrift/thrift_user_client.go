package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/services/thrift_gen/services"

	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	client := getClient(1661, false)
	ctx := context.Background()
	req := services.UserRequest{Name: "moye"}

	resp, err := client.GetUser(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[THRIFT_CLIENT]:\n %v\n", resp)
}

func getClient(port int, secure bool) *services.UserServiceClient {
	addr := fmt.Sprintf(":%d", port)
	var transport thrift.TTransport
	var err error

	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		log.Fatal(err)
	}

	transportFactory := thrift.NewTTransportFactory()
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		log.Fatal(err)
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		log.Fatal(err)
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := thrift.NewTStandardClient(iprot, oprot)
	return services.NewUserServiceClient(client)
}
