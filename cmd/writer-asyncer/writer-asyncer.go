package main

import (
	"write-async/internal/pkg/server"
	"write-async/internal/pkg/storage/inmemory"
)

func main() {
	db := inmemory.NewDatabase("test.txt")
	s := server.NewServer(db)
	s.Run()
}
