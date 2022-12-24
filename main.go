package main

import "chess/backend"

func main() {
	server := backend.NewServer("", ":8081")
	server.Init()
	server.Run()
}
