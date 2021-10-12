package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// ***Basics***
	// Step 1: Listen
	// Step 2: Accept
	// Step 3: Read/Write From That Connection

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
	}
	defer conn.Close()
	// Remeber Always To Close The File

	// We Cant Get Here Because Its A Open Stream Connections So Never Close Connection And Never Print The Code Above
	fmt.Println("We Dont Get Here")
}
