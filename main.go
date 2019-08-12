package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/hackercompany/StarTrek/cache"
	"github.com/hackercompany/StarTrek/codex"
)

func init() {
	// Initialising time and filename for log
	log.SetFlags(log.LstdFlags | log.Llongfile)

	flag.Parse()
}

func main() {

	// Joining all args into a string with " "
	var name = strings.Join(flag.Args(), " ")

	// Validate the final input name
	valid := validateInput(name)

	// Exit if validation failes
	if !valid {
		fmt.Println("Please enter a valid name")
		fmt.Println("Eg: Uhura")
	}

	// Initialize translator
	k := codex.NewKlingon()

	// Get hex translation string for given name
	hexTranslation, err := k.Translate(name)

	// Raise error if name cannot be translated
	if err != nil {
		fmt.Println("Cannot translate name", name)
		fmt.Println(err.Error())
		return
	}

	// Finding the species if translation raises no error
	character, err := cache.GetCharacterByName(name)

	if err != nil {
		fmt.Println(hexTranslation)
		fmt.Println(err.Error())
		return
	}

	fmt.Println(hexTranslation)

	// Printing the first species returned for the character
	fmt.Println(character.Species[0])
}

func validateInput(param string) bool {

	if param == "" {
		return false
	}

	return true
}
