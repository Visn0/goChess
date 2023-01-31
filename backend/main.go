package main

import "chess/server"

func main() {
	s := server.NewServer("", ":8081")
	s.Run()
}
