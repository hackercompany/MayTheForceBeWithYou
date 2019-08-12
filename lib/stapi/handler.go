package stapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/hackercompany/StarTrek/constants"
	libhttp "github.com/hackercompany/StarTrek/lib/http_handler"
	"github.com/hackercompany/StarTrek/logger"
)

// SearchCharacterInfoByName searches for the character
// details and its species using the name of the character
// if character or the species is not found
// specific errors are raised
func SearchCharacterInfoByName(name string) (uid string, species []string, err error) {

	character := stCharacter{Name: name}
	err = character.getCharacterUid()

	if err != nil {
		return "", nil, errors.New(constants.CHARACTER_NOT_FOUND)
	}

	err = character.getCharacterSpecies()

	if err != nil {
		return "", nil, errors.New(constants.SPECIES_NOT_FOUND)
	}

	species = make([]string, 0)

	for _, specie := range character.Species {
		species = append(species, specie.Name)
	}

	return character.Uid, species, nil
}

// Private function to find character uid.
func (st *stCharacter) getCharacterUid() error {

	// URL to get character Uid details
	uidURL := constants.StURL + constants.CHARACTER_SUB_URL

	name := strings.ToLower(strings.Trim(st.Name, " "))

	// Form the payload for the POST call
	form := url.Values{}
	form.Add("name", name)
	form.Add("title", name)

	characterBytes, err := libhttp.CallRest(uidURL, http.MethodPost, form.Encode())

	// Raising custom character based error aggrigating
	// all HTTP errors
	if err != nil {
		logger.Print("getCharacterUid", err.Error())
		return errors.New(constants.CHARACTER_NOT_FOUND)
	}

	var characterInfo = struct {
		Page struct {
			Total int `json:"totalElements"`
		}
		Characters []stCharacter `json:"characters"`
	}{}

	json.Unmarshal(characterBytes, &characterInfo)

	if len(characterInfo.Characters) == 0 {
		logger.Print("getCharacterUid", constants.CHARACTER_NOT_FOUND)
		return errors.New(constants.CHARACTER_NOT_FOUND)
	}

	// Multiple characters might be returned by the API
	// Selecting the most significant at 0th position
	st.Uid = characterInfo.Characters[0].Uid

	return nil
}

// For a given uid this tried if stapi has a species
func (st *stCharacter) getCharacterSpecies() error {
	// URL for finding the species
	speciesURL := constants.StURL + constants.SPECIES_SUB_URL + "?uid=" + st.Uid

	speciesBytes, err := libhttp.CallRest(speciesURL, http.MethodGet, "")

	// Raising custom error aggrigating all HTTP errors
	if err != nil {
		logger.Print("getCharacterSpecies", err.Error())
		return errors.New(constants.SPECIES_NOT_FOUND)
	}

	var speciesInfo = struct {
		Character stCharacter `json:"character"`
	}{}

	json.Unmarshal(speciesBytes, &speciesInfo)

	// if no species found raise error
	if len(speciesInfo.Character.Species) == 0 {
		logger.Print("getCharacterSpecies", constants.SPECIES_NOT_FOUND)
		return errors.New(constants.SPECIES_NOT_FOUND)
	}

	st.Species = speciesInfo.Character.Species

	return nil
}
