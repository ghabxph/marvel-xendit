package marvel

import (
	"strconv"
	"github.com/ghabxph/marvel-xendit/internal/live"
)

type Memorydb_impl interface {
	CreateCharacter(id int, name string, description string)
	GetCharacters(page int) string
	GetCharacter(id int) (string, bool)
}

type marvel struct {
	db Memorydb_impl
}

var instance *marvel

func GetInstance(db ...Memorydb_impl) *marvel {
	if instance == nil {
		instance = &marvel{db:db[0]}
		live.GetInstance(db[0])
	}
	return instance
}

// Get all marvel characters (max of 100)
func (m *marvel) GetAllCharacters(page ...int) string {
	if len(page) == 0 {
		return m.db.GetCharacters(1)
	}
	return m.db.GetCharacters(page[0])
}

// Get a marvel character by ID
func (m *marvel) GetCharacter(id_str string) string {
	// Convert string int to integer.
	// If parameter is not valid int, we will return an error.
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return err.Error()
	}

	// Check if the character exists in cache
	character, exists := m.db.GetCharacter(id)
	if exists {
		// Returns the character stored in cache
		return character
	}
	// Returns the character straight from Marvel and caches it
	return live.GetInstance().GetCharacter(id_str)
}
