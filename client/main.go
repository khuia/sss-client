package main

//"fmt"

var c Config

func main() {
	getConfig("./config.json", &c)
	socks5Start(&c)
}
