package main

import (
	"github.com/ghabxph/marvel-xendit/internal/memorydb"
	_gateway "github.com/ghabxph/marvel-xendit/internal/gateway"
)

func main() {

	// Create memorydb instance
	db := memorydb.GetInstance()

	// Create gateway instance
	gateway := _gateway.GetInstance(db)

	// Initialize fiber
	gateway.Fiber()

	// Start the server
	gateway.Serve()
}
