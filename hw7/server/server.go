package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

type Server struct {
	listener net.Listener
	address  string
}

func NewServer(address string) *Server {
	s := new(Server)
	s.address = address
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal(err)
	}
	s.listener = listener

	return s
}

func (s *Server) Serve() error {
	log.Println("[TCP] Server running on", s.address)
	defer s.listener.Close()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		go func() {
			s.handler(conn)
		}()
	}
}

func (s *Server) handler(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		req := scanner.Text()
		log.Println("Request:", req)

		value, err := strconv.Atoi(req)
		if err != nil {
			conn.Write([]byte("Bad request: expected integer\n"))
			continue
		}

		time.Sleep(200 * time.Millisecond)
		response := value * value
		conn.Write([]byte(fmt.Sprintf("%d\n", response)))
	}

	if err := scanner.Err(); err != nil {
		log.Println("error:", err)
	}
}
