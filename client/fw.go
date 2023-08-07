package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// 普通转发
func Forward(dst io.WriteCloser, src io.ReadCloser) {
	defer dst.Close()
	defer src.Close()
	fmt.Printf("transferring data from %s to %s\n", src.(net.Conn).RemoteAddr(), dst.(net.Conn).RemoteAddr())

	srcData, err := io.ReadAll(src)
	if err != nil {
		fmt.Println("failed to read data:", err)
		return
	}

	if _, err := dst.Write(srcData); err != nil {
		fmt.Println("failed to write data:", err)
		return
	}
}

// 进阶转发
func advForward(des net.Conn, src net.Conn) {
	defer des.Close()
	defer src.Close()
	buf := make([]byte, 4096)
	for {
		n, err := src.Read(buf)
		fmt.Println("4096读取的数据为", buf[:n])
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		if _, err := des.Write(buf[:n]); err != nil {
			log.Println(err)
			break
		}
	}
}

func adv_VToO_Forward(des net.Conn, src net.Conn, config *Config) {

	defer des.Close()
	defer src.Close()

	var pack Packet
	buf := make([]byte, 10240)

	for {
		n, err := src.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		//加密
		cipTxt, Nonce, err := Encrypt(config.Key, buf[:n])
		if err != nil {
			fmt.Println("encrypt error:", err)
		}
		pack.data = cipTxt
		pack.nonce = Nonce

		handData := pack.pack()
		des.Write(handData)
	}

}

func adv_OToV_Forward(des net.Conn, src net.Conn, c *Config) {
	cipTxt := unpack(src)
	srcTxt, err := Decrypt(c.Key, cipTxt)
	if err != nil {
		fmt.Println("decrypt error", err)
	}
	des.Write(srcTxt)

}
