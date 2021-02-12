package marvel

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"github.com/ghabxph/marvel-xendit/internal/live"
	"github.com/ghabxph/marvel-xendit/internal/memorydb"
	"github.com/ghabxph/marvel-xendit/internal/testutils"
	"github.com/gofiber/fiber/v2/utils"
)

func TestMarvel(t *testing.T) {

	// Disable logging
	log.SetOutput(ioutil.Discard)

	// Set config path
	os.Setenv(live.CONFIG_PATH_KEY, "../../config.yaml")

	// Create memorydb instance
	db := memorydb.GetInstance()

	// Prepare dataset
	testutils.PrepareDataset(db)

	// Create marvel instance
	marvel := GetInstance(db)


	t.Run("Get all characters", func(t *testing.T) {

		// Get all characters in page 1
		chars := marvel.GetAllCharacters("1")

		// Do we get all characters?
		utils.AssertEqual(t, testutils.GetTestCharacters(), chars)
	})

	t.Run("Get a character", func(t *testing.T) {
		// Get a character
		char := marvel.GetCharacter("1009146")

		// Do we get the character?
		utils.AssertEqual(t, testutils.GetTestCharacter(), char)
	})
}
