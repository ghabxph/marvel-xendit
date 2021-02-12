package scraper

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"github.com/ghabxph/marvel-xendit/internal/live"
	"github.com/ghabxph/marvel-xendit/internal/memorydb"
	"github.com/gofiber/fiber/v2/utils"
)

const e_characters string = "[1009144, 1009146, 1009148, 1009149, 1009150, 1009151, 1009152, 1009153, 1009154, 1009156, 1009158, 1009159, 1009160, 1009161, 1009162, 1009163, 1009164, 1009165, 1009168, 1009169, 1009170, 1009171, 1009172, 1009173, 1009174, 1009175, 1009176, 1009177, 1009178, 1009179, 1009240, 1009329, 1009346, 1009435, 1009497, 1009550, 1009567, 1009596, 1009740, 1010336, 1010354, 1010370, 1010672, 1010673, 1010674, 1010686, 1010699, 1010718, 1010748, 1010755, 1010773, 1010784, 1010801, 1010802, 1010827, 1010835, 1010836, 1010846, 1010866, 1010870, 1010903, 1010905, 1010906, 1010908, 1010909, 1011012, 1011031, 1011120, 1011136, 1011137, 1011164, 1011170, 1011175, 1011176, 1011194, 1011198, 1011208, 1011214, 1011227, 1011253, 1011266, 1011275, 1011297, 1011298, 1011301, 1011324, 1011334, 1011338, 1011354, 1011361, 1011382, 1011396, 1011766, 1014990, 1015239, 1016823, 1016824, 1017100, 1017438, 1017574]"
const e_character string = `{
  "id": 1010909,
  "name": "Beast (Earth-311)",
  "description": ""
}`

type mockHttp struct{}
func (m *mockHttp) Get(url string) string {
	file, _ := ioutil.ReadFile("test_characters.json")
	return string(file)
}

func TestScraper(t *testing.T) {

	// Disable logging
	log.SetOutput(ioutil.Discard)

	// Set config path
	os.Setenv(live.CONFIG_PATH_KEY, "../../config.yaml")

	// Creates MemoryDB instance
	db := memorydb.GetInstance()

	// Creates Scraper Instance
	scraper := GetInstance(db)

	// Creates mock http (from json file instead)
	scraper.MockHttpGet(&mockHttp{})

	// Runs the scraper (creates new thread)
	scraper.RunScraper()

	t.Run("Get and store characters in Scraper", func(t *testing.T) {

		// Scrapes for page 5
		scraper.Scrape(5)

		// Waits for scraper to end
		scraper.ScrapeNext()
	})

	t.Run("Check characters (page 1) in mock MemoryDB", func(t *testing.T) {

		// Get the characters in page 1
		// TODO: Page parameter is still not yet functional
		// Right now, it will return all items you've collected.
		// Maybe one test point is to make sure that page 1 must only return 100 items.
		characters := db.GetCharacters(1)

		// Check if all character exist.
		utils.AssertEqual(t, e_characters, characters)
	})

	t.Run("Check a character in mock MemoryDB", func(t *testing.T) {
		character, exists := db.GetCharacter(1010909)

		// Character must exist
		utils.AssertEqual(t, true, exists)

		// Character should be this one.
		utils.AssertEqual(t, e_character, character)
	})
}
