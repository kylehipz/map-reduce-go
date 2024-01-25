package master

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type JoinArgs struct {
	addr string // Address of the worker
}

type Pool struct{}

func (p *Pool) Join(args *JoinArgs, reply *string) error {
	*reply = "success"

	return nil
}

// Initialize master node
func Init(addr string, mWorkers, rWorkers int) {
	// Setup an rpc server using the address
	// Keep track of idle worker nodes
	// Assign worker node a task (map/reduce)
	p := Pool{}

	rpc.Register(&p)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Listen error:", err)
	}

	http.Serve(listener, nil)
}
