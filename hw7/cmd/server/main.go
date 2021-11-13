package main

import (
	"github.com/one-lab-go/hw7/server"
	"os"
	"os/signal"
	"syscall"
)

const Port = ":8081"

func main() {
	s := server.NewServer(Port)
	c := make(chan os.Signal, 1)

	go func() {
		err := s.Serve()
		if err != nil {
			panic(err)
		}
	}()

	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
}