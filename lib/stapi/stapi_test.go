package stapi

import (
	"github.com/hackercompany/StarTrek/constants"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterNotFound(t *testing.T) {
	uid, _, err := SearchCharacterInfoByName("Ashmeet")
	assert.Equal(t, uid, "")
	assert.Equal(t, err.Error(), constants.CHARACTER_NOT_FOUND)
}

func TestCharacterFound(t *testing.T) {
	uid, _, err := SearchCharacterInfoByName("Uhura")
	assert.Equal(t, uid, "CHMA0000115364")
	assert.Nil(t, err)
}

func TestUidSuccess(t *testing.T) {
	character := stCharacter{Name: "Pechetti"}
	err := character.getCharacterUid()
	assert.Equal(t, character.Uid, "CHMA0000021696")
	assert.Nil(t, err)
}

func TestUidFail(t *testing.T) {
	character := stCharacter{Name: "Jexia"}
	err := character.getCharacterUid()
	assert.Equal(t, character.Uid, "")
	assert.Equal(t, err.Error(), constants.CHARACTER_NOT_FOUND)
}

func TestSpeciesSuccess(t *testing.T) {
	character := stCharacter{Name: "T. Virts", Uid: "CHMA0000101321"}
	err := character.getCharacterSpecies()
	assert.Nil(t, err)
	assert.Equal(t, character.Species[0].Name, "Human")
	assert.Equal(t, character.Species[0].Uid, "SPMA0000026314")
}

func TestSpeciesFailure(t *testing.T) {
	character := stCharacter{Name: "Pomet", Uid: "CHMA0000028502"}
	err := character.getCharacterSpecies()
	assert.EqualError(t, err, constants.SPECIES_NOT_FOUND)
}
