package memorydb

import (
	"strconv"
	"strings"
	"sort"
)

type memorydb struct{}
var instance *memorydb

type character struct {
	id int
	name string
	description string
}

// Map of characters
var characters map[int]*character

// Get memorydb instance
func GetInstance() *memorydb {
	if instance == nil {
		instance = &memorydb{}
	}
	return instance
}

// Create new character and store in memory
func (m *memorydb) CreateCharacter(id int, name string, description string) {

	// Create map of characters if there's none yet
	if characters == nil {
		characters = make(map[int]*character)
	}

	// Store new character in memory
	characters[id] = &character{id:id, name:name, description:description}

	// Store new character in filesystem
	characters[id].store()
}

// Create new character
func (m *memorydb) GetCharacter(id int) (character string, exists bool) {
	defer func(character *string, exists *bool) {
		if r := recover(); r != nil {
			*character = "{\"error\": \"Character does not exist\"}"
			*exists = false
		}
	}(&character, &exists)

	return characters[id].getString(), true
}

func (m *memorydb) GetCharacters(page int) string {
	// Slice of int keys
	keys_i := make([]int, 0, len(characters))

	// Slice of converted int key strings
	keys := make([]string, 0, len(characters))

	// There's no convenient way to convert integer keys of a map.
	// If there is, I wouldn't resort to loop. I hope there's a O(1)
	// solution here...
	for k := range characters {
		keys_i = append(keys_i, k)
	}

	// Sort the keys
	sort.Ints(keys_i)
	for _, v := range keys_i {
		keys = append(keys, strconv.Itoa(v))
	}

	return "[" + strings.Join(keys, ", ") + "]"
}

// Get json string of character
func (c *character) getString() string {
	return `{
  "id": ` + strconv.Itoa(c.id) + `,
  "name": "` + c.name + `",
  "description": "` + c.description + `"
}`
}

// Store character in filesystem
func (c *character) store() {
	// TODO: Store character in filesystem
}
