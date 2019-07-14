package services

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/rockdragon/daily-practice/code/golang/mod_proj/services/thrift_gen/services"
)

func ServeThrift(port int, secure bool) {
	addr := fmt.Sprintf(":%d", port)
	var transport thrift.TServerTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		} else {
			log.Fatal(err)
		}
		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTServerSocket(addr)
	}

	if err != nil {
		log.Fatal(err)
	}

	handler := UserServiceHandler{}
	processor := services.NewUserServiceProcessor(&handler)
	server := thrift.NewTSimpleServer4(processor, transport,
		thrift.NewTTransportFactory(),
		thrift.NewTBinaryProtocolFactoryDefault())

	log.Printf("[THRIFT] services.UserServer is listening on port: %d\n", port)

	log.Fatal(server.Serve())
}
