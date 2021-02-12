package scraper

import (
	"github.com/ghabxph/marvel-xendit/internal/live"
	"log"
	"strconv"
	"time"
)

type live_impl interface {
	GetCharacters(offset string) int
	MockHttpGet(mock interface{})
}

type scraper struct {
	db                interface{}
	scraper_running   bool
	offset            chan int
	live              live_impl
	scraping_done     chan struct{}
	total_characters  chan int
}

var instance *scraper

// Gets scraper singleton instance
func GetInstance(db ...interface{}) *scraper {
	if instance == nil {
		instance = &scraper{db: db[0], offset: make(chan int, 20), live: live.GetInstance(db[0]), scraping_done: make(chan struct{}), total_characters: make(chan int)}
	}

	return instance
}

// Initialize the scraper in the background
func (s *scraper) Start() {

	// Ticker that runs the scraper for everyday
	scrape_ticker := time.NewTicker(24 * time.Hour)

	// Sleep for 5 seconds.
	time.Sleep(5 * time.Second)

	// Starts the scraper
	s.RunScraper()

	// Keeps the scraper to continue
	go func() { for { s.ScrapeNext() } }()

	// Scraper that runs every day
	go func() {

		// Infinite loop
		for {

			// 1. Gets page 1
			log.Println("Getting page #1 aiming to get total number of characters.")
			s.Scrape(1)

			// 2. Gets total characters
			total_characters := <-s.total_characters

			// 3. Scrape!
			log.Println("There's a change in total number of characters. We are now going to begin the scraper.")

			for i := 2; i <= total_characters/100+1; i++ {

				// Scraping page #x
				s.Scrape(i)
			}

			// Notify
			log.Println("Scraper is done. Next scraping schedule will be tommorow same time this program is executed")
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

			// Notify
			log.Println("Scraping offset #" + strconv.Itoa(offset) + " to Marvel server...")

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

			// Notify
			log.Println("Scraping offset #" + strconv.Itoa(offset) + " done.")

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

	// Notify
	log.Println("Sending offset #" + strconv.Itoa(offset) + " to the queue.")

	// Sends the offset to offset channel for processing
	s.offset <- offset
}

// Moves to next scrape item
func (s *scraper) ScrapeNext() {

	// Clears the channel, thus scraper goroutine can continue.
	<-s.scraping_done

	// Notify
	log.Println("Moving to next item...")
}

// Mocks HttpGet (For testing)
func (s *scraper) MockHttpGet(mock interface{}) {
	s.live.MockHttpGet(mock)
}
