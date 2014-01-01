package main

import "fmt"
import "os"
import "net"
import "net/rpc"
import "nettest/mymath"

func main() {

	otherend := new(OtherEnd)
	rpc.Register(otherend)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing", err)
	}
	args := os.Args[1:]
	
	for idx, host := range args {
		fmt.Println(idx, host)
		_, err := net.Dial("udp", host)
		if err != nil {
			fmt.Println(" Caught ", err)
			continue
		}
		fmt.Println(host)
	}
}
