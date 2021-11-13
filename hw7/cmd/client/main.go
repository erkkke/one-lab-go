package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		fmt.Println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()


	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Request: ")
		text, _ := reader.ReadString('\n')

		conn.Write([]byte(text))
		res, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Server response: " + res)
	}
}
