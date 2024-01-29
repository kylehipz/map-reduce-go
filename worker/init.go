package worker

import (
	"fmt"
	"log"
	"net"
)

func Init() {
	addr := "localhost:8000"

	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatal("Dial error:", err)
	}

	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)
		if err != nil {
			break
		}

		fmt.Println(string(buf))
	}
}
