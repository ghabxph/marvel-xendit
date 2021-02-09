package memorydb

import (
	"testing"
	"github.com/gofiber/fiber/v2/utils"
)

func TestMemoryDB(t *testing.T) {

	// MemoryDB Instance
	db := GetInstance()

	// Characters (dataset)
	characters := "[1009146]"
	character := `{
  "id": 1009146,
  "name": "Abomination (Emil blonsky)",
  "description": "Formerly known as Emil Blonsky, a spy of Soviet Yugoslavian origin working for the KGB, the Abomination gained his powers after receiving a dose of gamma radiation similar to that which transformed Bruce Banner into the incredible Hulk."
}`

	t.Run("Create a character in memory", func(t *testing.T) {
		db.CreateCharacter(
			1009146,
			"Abomination (Emil blonsky)",
			"Formerly known as Emil Blonsky, a spy of Soviet Yugoslavian origin working for the KGB, the Abomination gained his powers after receiving a dose of gamma radiation similar to that which transformed Bruce Banner into the incredible Hulk.",
		)
	})

	t.Run("Get a character in memory", func(t *testing.T) {
		char, _ := db.GetCharacter(1009146)
		utils.AssertEqual(t, character, char)
	})

	t.Run("Get all characters in page 1 in memory", func(t *testing.T) {
		chars := db.GetCharacters(1)
		utils.AssertEqual(t, characters, chars)
	})
}
