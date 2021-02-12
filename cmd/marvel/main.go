package main

import (
	"os"
	"github.com/ghabxph/marvel-xendit/internal/memorydb"
	_gateway "github.com/ghabxph/marvel-xendit/internal/gateway"
	//_scraper "github.com/ghabxph/marvel-xendit/internal/scraper"
)

func main() {

	// Well basically, uhm.. yeah.
	os.Setenv("MARVEL_XENDIT_PATH", "config.yaml")

	// Create memorydb instance
	db := memorydb.GetInstance()

	// Load characters from file system
	db.Load()

	// Create gateway instance
	gateway := _gateway.GetInstance(db)

	// Creates scraper instance
	//scraper := _scraper.GetInstance(db)

	// Starts the scraper in the background
	//go scraper.Start()

	// Initialize fiber
	gateway.Fiber()

	// Start the server
	gateway.Serve()
}
