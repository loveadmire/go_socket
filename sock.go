package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("hello world")
	addr := "127.0.0.1:8080"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("访问本地:", conn.RemoteAddr().String())
	fmt.Println("my port is \n", conn.LocalAddr())
	n, err := conn.Write([]byte("hello it is me!\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("send data to server \n")
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if (err != nil) && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
	conn.Close()

}
