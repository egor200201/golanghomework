package main

import (
	"bufio" // buffers for reading
	"fmt"   // output into terminal
	"log"   // logging pkg
	"net"   // main pkg for networking
	"strconv"
	"strings"
)

func main() {
	// server port number
	const port = 8081

	fmt.Printf("Launching server on port: %d \n\n", port)

	// create listener
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		// log error and exit with error code
		log.Fatal(err)
	}

	// endless loop for listening connections
	for {
		// accept new connection
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// serve connection
		go func(c net.Conn) {
			// call close function after serving
			defer c.Close()

			// read string from client(string should end with '\n' symbol)
			text, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				log.Print(err)
				return
			}

			fmt.Printf("Message received: %s \n", text)

			text = text[:len(text)-1]
			txt := strings.TrimSpace(text)
			res, err := strconv.Atoi(txt)
			var returntext string
			if err != nil {
				returntext = strings.ToUpper(text)
			} else {
				returntext = strconv.Itoa(res * 2)
			}
			// write message back(appends with '\n', because it deleted in previous step)
			_, err = c.Write([]byte(returntext + "\n"))
			if err != nil {
				log.Print(err)
				return
			}
		}(conn)
	}
}
