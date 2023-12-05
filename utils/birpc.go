package utils

import (
	"addons/api"
	"log"
	"net"
	"net/http/httptest"
	"sync"

	"github.com/cgrates/birpc"
)

var (
	httpServerAddr string
	httpOnce       sync.Once
)

func listenTCP() (net.Listener, string) {
	l, e := net.Listen("tcp", "127.0.0.1:2020") // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	return l, l.Addr().String()
}

func StartServerBiRPC() {
	birpc.Register(new(api.AddonS))
	// birpc.Register(new(Embed))
	// birpc.RegisterName("net.rpc.Arith", new(Arith))
	// birpc.Register(BuiltinTypes{})

	var l net.Listener
	l, serverAddr := listenTCP()
	log.Println("Test RPC server listening on", serverAddr)
	go birpc.Accept(l)

	birpc.HandleHTTP()
	httpOnce.Do(startHttpServer)
}

func startHttpServer() {
	server := httptest.NewServer(nil)
	httpServerAddr = server.Listener.Addr().String()
	log.Println("Test HTTP RPC server listening on", httpServerAddr)
}
