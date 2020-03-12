package test

import (
	"github.com/cloudflare/cfssl/log"
	"net"
	"net/rpc"
)

type HelloService struct {}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request

	return nil
}

func RPCService() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":9897")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error: ", err)
	}

	rpc.ServeConn(conn)
}

func RPCClient() {
	//client, err := rpc.Dial("tcp", "localhost:9897")
	//if err != nil {
	//	log.Fatal("RPC dial error: ", err)
	//}
	//
	//rpc.ServeRequest()
	//jsonrpc.NewServerCodec()

}
