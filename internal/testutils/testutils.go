package testutils


// Characters (dataset)
const characters string = "[1009146, 1009148, 1010699, 1011334, 1016823, 1017100]"
const character string = `{
  "id": 1009146,
  "name": "Abomination (Emil blonsky)",
  "description": "Formerly known as Emil Blonsky, a spy of Soviet Yugoslavian origin working for the KGB, the Abomination gained his powers after receiving a dose of gamma radiation similar to that which transformed Bruce Banner into the incredible Hulk."
}`

type memorydb_impl interface {
	CreateCharacter(id int, name string, description string)
}

func PrepareDataset(db memorydb_impl) {
	db.CreateCharacter(1011334, "", "")
	db.CreateCharacter(1017100, "", "")
	db.CreateCharacter(1009146,
		"Abomination (Emil blonsky)",
		"Formerly known as Emil Blonsky, a spy of Soviet Yugoslavian origin working for the KGB, the Abomination gained his powers after receiving a dose of gamma radiation similar to that which transformed Bruce Banner into the incredible Hulk.",
	)
	db.CreateCharacter(1010699, "", "")
	db.CreateCharacter(1016823, "", "")
	db.CreateCharacter(1009148, "", "")
}

func GetTestCharacters() string {
	return characters
}

func GetTestCharacter() string {
	return character
}
