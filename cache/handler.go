package cache

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/hackercompany/MayTheForceBeWithYou/constants"
	"github.com/hackercompany/MayTheForceBeWithYou/lib/stapi"
	"github.com/hackercompany/MayTheForceBeWithYou/logger"
)

var cachedCharacters map[string]Character
var externalCharacterSearchViaName libInterface

func init() {

	// Select the external api to be called
	// This can be made configurable via command line args as well
	switch constants.API_TAG {
	case "stapi":
		externalCharacterSearchViaName = stapi.SearchCharacterInfoByName
	}

	// Reading data from json dump file
	if data, err := ioutil.ReadFile(constants.DATA_FILE); err == nil {
		json.Unmarshal(data, &cachedCharacters)
	} else {
		cachedCharacters = make(map[string]Character)
	}
}

// GetCharacterByName searches character info by name
func GetCharacterByName(name string) (c Character, err error) {

	// Lookup in cache before going forward
	if c, ok := cachedCharacters[name]; ok {
		return c, nil
	}

	// If data is not found in cache, API is hit to get information
	// about the character
	uid, species, err := externalCharacterSearchViaName(name)

	// There can be multiple reason for this error.
	// From invalid name to API not responding with any species
	if err != nil {
		return c, errors.New(constants.SPECIES_NOT_FOUND)
	}

	// Global cache is updated with current new information
	cachedCharacters[name] = Character{Uid: uid, Species: species}

	updateCache()

	return cachedCharacters[name], nil
}

// Update the file cache after every successful read
func updateCache() error {

	data, err := json.Marshal(cachedCharacters)

	if err != nil {
		logger.Print("updateCache", err.Error())
		return err
	}

	// Ovewriting the data file present in the directory
	if err = ioutil.WriteFile(constants.DATA_FILE, data, 0645); err != nil {
		logger.Print("updateCache", err.Error())
		return err
	}

	return nil
}
