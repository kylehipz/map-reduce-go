package master

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"
)

// Initialize master node
func Init(directory string) {
	// Requirements:
	// Setup an tcp server using the address
	// Keep track of idle worker nodes
	// Assign worker node a task (map/reduce)

	// Flow
	// Background Task:
	// Accept tcp connections
	// 1 connection = 1 worker
	// Create a worker struct from the connection
	// Worker struct: address, connection to write to,
	// state (idle, in-progress, done), task (map, reduce, none)

	// Get files from directory
	// Assign m workers with map
	// Assign r workers with map

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln("Listen error:", err)
	}

	workers := []net.Conn{}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Connection error", err)
				continue
			}

			workers = append(workers, conn)
		}
	}()

	dirEnts, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalln("Cannot read file", err)
	}

	for _, entry := range dirEnts {
		absPath := path.Join(directory, entry.Name())
		fmt.Println(absPath)
	}
}
