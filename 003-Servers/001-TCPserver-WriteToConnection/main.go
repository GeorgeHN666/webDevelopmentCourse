package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Step 1 Listen:
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close() // We Have To Close The Listener

	// Step 2 : LOOP
	// We Loop And If Somebosy Calls We Accept The Connection
	for {
		connection, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}

		io.WriteString(connection, "\nHello From TCP Server\n")
		fmt.Fprintln(connection, "How Is Your Day?")
		fmt.Fprintf(connection, "%v", "Well, I Hope!")

		connection.Close()
	}
}
