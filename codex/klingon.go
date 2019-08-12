package codex

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/hackercompany/MayTheForceBeWithYou/constants"
)

// Master codex for klingon
var klingonMaster map[string]string

// Data flow object
type klingon struct {
	lookupTable *map[string]string
}

func init() {
	// Load klingon dictionary from file
	if data, err := ioutil.ReadFile(constants.KLINGON_DATA); err == nil {
		json.Unmarshal(data, &klingonMaster)
	} else {
		// If file is not found then panic
		panic(constants.DICTIONARY_CORRUPTED)

	}
}

// NewKlingon initialises klingon object with a reference
// to the master lookup table
func NewKlingon() *klingon {
	return &klingon{lookupTable: &klingonMaster}
}

// Translate tries to translate character by character
// If a character is not found in the lookup table
// error is raised
func (k *klingon) Translate(value string) (string, error) {

	// Raise error if empty value is passed
	if value == "" {
		return "", errors.New(constants.INVALID_INPUT)
	}

	translated := []string{}

	length := len(value)

	cursor := 0

	// Loop until the end of string
	for cursor < length {

		// Looking for specific case "tlh"
		// if found no further evaluation  is done and loop jumps 3 steps
		if cursor+2 < length {
			translatedChar, ok := (*k.lookupTable)[strings.ToLower(value[cursor:cursor+3])]

			if ok {
				translated = append(translated, translatedChar)
				cursor += 3
				continue
			}
		}

		// Looking for cases "ch", "gh" and "ng"
		// Jumps 2 steps if found
		if cursor+1 < length {

			translatedChar, ok := (*k.lookupTable)[strings.ToLower(value[cursor:cursor+2])]

			if ok {
				translated = append(translated, translatedChar)
				cursor += 2
				continue
			}
		}

		// Looking for specific case Q
		// If not found continues to lower case
		currentChar := string(value[cursor])
		translatedChar, ok := (*k.lookupTable)[currentChar]

		if !ok {
			translatedChar, ok = (*k.lookupTable)[strings.ToLower(currentChar)]

			// If none of the cases are true then the character cannot be translated
			// to klingon and an error is raised
			if !ok {
				return "", errors.New(constants.INVALID_INPUT)
			}
		}

		// Appending the final cases output to translate array
		translated = append(translated, translatedChar)

		cursor++

	}

	// Join all hex reprensentation of words with " " separator
	return strings.Join(translated, " "), nil

}
