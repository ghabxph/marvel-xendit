package scraper

import (
	"github.com/ghabxph/marvel-xendit/internal/live"
	"log"
	"strconv"
	"time"
)

type memorydb_impl interface {
	TotalCharacterChanged(total int) bool
}

type live_impl interface {
	GetCharacters(offset string) int
	MockHttpGet(mock interface{})
}

type scraper struct {
	db               memorydb_impl
	scraper_running  bool
	offset           chan int
	live             live_impl
	scraping_done    chan struct{}
	total_characters chan int
}

var instance *scraper

// Gets scraper singleton instance
func GetInstance(db ...interface{}) *scraper {
	if instance == nil {
		instance = &scraper{db: db[0].(memorydb_impl), offset: make(chan int, 10), live: live.GetInstance(db[0]), scraping_done: make(chan struct{}), total_characters: make(chan int)}
	}

	return instance
}

// Initialize the scraper in the background
func (s *scraper) Start() {

	// Ticker that runs the scraper for every 1 day
	scrape_ticker := time.NewTicker(1 * time.Minute)

	// Starts the scraper
	s.RunScraper()

	// Keeps the scraper to continue
	go func() { s.ScrapeNext() }()

	// Scraper that runs every day
	go func() {

		// Infinite loop
		for {

			// 1. Gets page 1
			log.Println("Getting page #1 aiming to get total number of characters.")
			s.Scrape(1)

			// 2. Gets total characters
			total_characters := <-s.total_characters
			log.Println("Total number of characters:", total_characters)

			// 3. Checks if there's change in total number of characters
			if s.db.TotalCharacterChanged(total_characters) {

				// 4. Scrape!
				log.Println("There's a change in total number of characters. We are now going to begin the scraper.")
				log.Println("Scraping done.")
			}

			// 5. Waits for the next ticker...
			log.Println("Waits for the next ticker...")
			<-scrape_ticker.C
		}
	}()
}

// Runs the scraper goroutine
func (s *scraper) RunScraper() {

	// Exits the method if scraper is already running
	if s.scraper_running {
		return
	}

	go func() {

		// Infinite loop
		for {

			// Gets the offset.
			offset := <-s.offset

			// If offset is 0, we will send the total number of characters to total_characters channel
			if offset == 0 {

				// Gets the characters from Marvel
				// The method will also cache it automatically.
				s.total_characters <- s.live.GetCharacters(strconv.Itoa(offset))
			} else {

				// Gets the characters from Marvel
				// The method will also cache it automatically.
				s.live.GetCharacters(strconv.Itoa(offset))
			}

			// Notify that scraping is done
			s.scraping_done <- struct{}{}
		}
	}()

	s.scraper_running = true
}

// Scrapes a target page
func (s *scraper) Scrape(page int) {

	// Gets the offset
	offset := (page - 1) * 100

	// Sends the offset to offset channel for processing
	s.offset <- offset
}

// Moves to next scrape item
func (s *scraper) ScrapeNext() {

	// Clears the channel, thus scraper goroutine can continue.
	<-s.scraping_done
}

// Mocks HttpGet (For testing)
func (s *scraper) MockHttpGet(mock interface{}) {
	s.live.MockHttpGet(mock)
}
