package main

import (
	"net"
)

func handleConnect(conn net.Conn, config *Config) {
	defer conn.Close()

	go socks5Start(config)

	sockConn, _ := net.Dial("tcp", config.Socks5Host+":"+config.Socks5Port)

	go adv_OToV_Forward(sockConn, conn, config)
	go adv_VToO_Forward(conn, sockConn, config)

}
