package main

import (
	"fmt"
	"net"
)

func start(config *Config) {

	fmt.Println(config)

	listener, err := net.Listen("tcp", net.JoinHostPort(config.LocalHost, config.LocalPort))
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	fmt.Printf("Listening on %s:%s...\n", config.LocalHost, config.LocalPort)
	fmt.Println(config)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnect(conn, config)

	}
}
