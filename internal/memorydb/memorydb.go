package memorydb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

var store_mode bool = true

type memorydb struct{}
var instance *memorydb

type character struct {
	Id int
	Name string
	Description string
}

// Map of characters
var characters map[int]*character
// Slice of character ids
var character_ids []int

// Get memorydb instance
func GetInstance() *memorydb {

	// Check if instance does not exist yet
	if instance == nil {

		// Create instance
		instance = &memorydb{}

		// Initialize map of characters
		characters = make(map[int]*character)

		// Initialize slice of characters ids
		character_ids = make([]int, 0)
	}

	return instance
}

// Create new character and store in memory
func (m *memorydb) CreateCharacter(id int, name string, description string) {

	// Add the new ID to slice if it does not exist yet on characters map
	if _, exists := characters[id]; exists == false {

		// Adds ID to slice
		character_ids = append(character_ids, id)

		// Sorts the keys every new entry
		sort.Ints(character_ids)
	}

	// Store new character in memory
	characters[id] = &character{Id:id, Name:name, Description:description}

	// Store new character in filesystem
	characters[id].store()
}

// Create new character
func (m *memorydb) GetCharacter(id int) (character interface{}, exists bool) {

	if character := characters[id]; character == nil {
		return map[string]string{ "error": "Character does not exist." }, false
	}

	return map[string]interface{}{
		"id": characters[id].Id,
		"name": characters[id].Name,
		"description": characters[id].Description,
	}, true
}

func (m *memorydb) GetCharacters(page int) ([]int, int) {

	// Count of items to show
	count := 500

	// Total characters
	total := len(character_ids)

	// Gets the offset
	offset := (page - 1) * count

	// Returns empty result if offset is larger or equal to capacity of slice
	if offset >= total { return make([]int, 0), 404 }

	// Adjust the value of total if new total is less or equal total
	if new_total := offset + count; new_total <= total { total = new_total }

	// Slice of converted int key strings
	keys := make([]int, 0)

	for i := offset; i < total; i++ {

		// Adds the character id to string slice
		keys = append(keys, character_ids[i])
	}

	return keys, 200
}

// Get json string of character
func (c *character) getString() string {
	return `{
  "id": ` + strconv.Itoa(c.Id) + `,
  "name": "` + c.Name + `",
  "description": "` + c.Description + `"
}`
}

// Store character in filesystem
func (c *character) store() {

	// Exit the function if store_mode is disabled
	if store_mode == false { return	}

	// Check if .characters folder exists
	if _, err := os.Stat(".characters"); os.IsNotExist(err) {

		// Creates character folder
		os.Mkdir(".characters", 0755)
	}

	// Notify
	log.Println("Saving", strconv.Itoa(c.Id), ":", c.Name, " to file system.")

	// Storing character data to file system
	ioutil.WriteFile(".characters/" + strconv.Itoa(c.Id) + ".json", []byte(c.getString()), 0644)
}

// Loads character from FS
func (m *memorydb) Load() {

	// Notify
	log.Println("Checks if .character folder exists")

	// Check if .characters folder exists
	if _, err := os.Stat(".characters"); os.IsNotExist(err) {

		// Notify
		log.Println(".characters folder does not exist. Program will now proceed.")

		return
	}

	// Notify
	log.Println("Folder exist. Loading all characters...")

	// Disable auto store mode
	store_mode = false

	filepath.Walk(".characters", func(path string, f os.FileInfo, err error) error {

		// Skip if path is the target folder.
		if path == ".characters" {
			return err
		}

		// Create character instance
		char := character{}

		// Read character file
		file, _ := ioutil.ReadFile(path)

		// Creates character instance from json
		json.Unmarshal(file, &char)

		// Notify
		log.Println("Loading: ", char)

		// Stores character info to RAM
		GetInstance().CreateCharacter(char.Id, char.Name, char.Description)

		return err
	})

	// Re-enable store mode
	store_mode = true
}
