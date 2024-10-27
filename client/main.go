package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		fmt.Println(msg)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal("error client = ", conn)
	}
	fmt.Println("conn = ", conn)

	fmt.Println("-------------")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')

	nameInput = nameInput[:len(nameInput)-1]
	fmt.Println("name Input = ", nameInput)
	fmt.Println("***********")

	go onMessage(conn)

	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')

		if err != nil {
			break
		}

		msg = fmt.Sprintf("%s %s\n", nameInput, msg[:len(msg)-1])
		conn.Write([]byte(msg))
	}

	conn.Close()

}
