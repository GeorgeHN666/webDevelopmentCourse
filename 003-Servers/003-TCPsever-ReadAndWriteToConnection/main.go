package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I Heard You Say : %s\n", ln)
	}
	defer conn.Close()

	// We Never Get Here
	// We Have An Open Stream Connection
	fmt.Println("Code Got Here")
}
