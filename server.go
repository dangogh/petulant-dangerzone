package main

//import "fmt"
import "log"
import "net"
import "net/rpc"
import "github.com/dangogh/petulant-dangerzone/mymath"

func main() {
	const msgLen = 1024
	rpc.Register(new(mymath.Quotient))

	// listen on the port
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listening...", err)
	}
	for {
		go rpc.Accept(listen)

	}
}
