package main

import (
	"bufio"
	"flag"
	"log"
	"net"
)

var localAddr *string = flag.String("l", "localhost:8080", "local address:port")
var remoteAddr *string = flag.String("r", "localhost:80", "remote address:port")

func main() {
	flag.Parse()

	log.Printf("Proxying TCP traffic from %v to %v\n", *localAddr, *remoteAddr)

	listener, err := net.Listen("tcp", *localAddr)
	check(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		check(err)

		go handleRequest(conn)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	// Make a buffer and read incoming data into it
	buf := make([]byte, 1024)
	length, err := conn.Read(buf)
	check(err)

	// Connect to remote server
	rConn, err := net.Dial("tcp", *remoteAddr)
	check(err)
	defer rConn.Close()
	log.Println("Received request from client")

	// Get remote server response and send it to client
	rConn.Write([]byte(string(buf[:length])))
	connbuf := bufio.NewReader(rConn)
	for {
		response, err := connbuf.ReadString('\n')
		check(err)
		if len(response) > 0 {
			conn.Write([]byte(string(response)))
			log.Println("Sent response to client:", string(response))
		}
		break
	}
	log.Println("Finished communicating with client")
}
