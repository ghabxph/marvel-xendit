package marvel

import (
	"github.com/ghabxph/marvel-xendit/internal/testutils"
	"github.com/ghabxph/marvel-xendit/internal/memorydb"
	"github.com/gofiber/fiber/v2/utils"
	"testing"
)

func TestMarvel(t *testing.T) {

	// Create memorydb instance
	db := memorydb.GetInstance()

	// Prepare dataset
	testutils.PrepareDataset(db)

	// Create marvel instance
	marvel := GetInstance(db)


	t.Run("Get all characters", func(t *testing.T) {
		// Get all characters
		chars := marvel.GetAllCharacters()

		// Do we get all characters?
		utils.AssertEqual(t, testutils.GetTestCharacters(), chars)
	})

	t.Run("Get a character", func(t *testing.T) {
		// Get a character
		char := marvel.GetCharacter("1009146")
		//char := marvel.GetCharacter("1010354")

		// Do we get the character?
		utils.AssertEqual(t, testutils.GetTestCharacter(), char)
	})
}
