package main

import "chess/server"

func main() {
	server := server.NewServer("", ":8081")
	server.Run()
}
