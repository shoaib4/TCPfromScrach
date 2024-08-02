package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Main started")
	listiner, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal("error while creating listiner")
	}
	conn, err := listiner.Accept()
	defer conn.Close()
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		log.Fatal("error in reading data")
	}
	//fmt.Println(buff[:n])
	fmt.Println("read size = ", n, "\n", string(buff[:n]))
	//time.Sleep(3 * time.Second)
	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello from the other side!!!\r\n"))
	if err != nil {
		log.Fatal("error in writing data")
	}

	//for {
	//}
}
