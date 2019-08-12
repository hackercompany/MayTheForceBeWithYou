package cache

import (
	"os"
	"testing"

	"github.com/hackercompany/MayForceBeWithYou/constants"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCharacterFileSuccess(t *testing.T) {
	cachedCharacters["uhura"] = Character{
		Uid:     "CHMA0000115364",
		Species: []string{"Human"},
	}
	assert.Nil(t, updateCache())

	os.Remove("data.txt")
}

func TestReadFromFileWithCache(t *testing.T) {
	cachedCharacters["uhura"] = Character{
		Uid:     "CHMA0000115364",
		Species: []string{"Human"},
	}
	assert.Nil(t, updateCache())

	character, err := GetCharacterByName("uhura")

	assert.Nil(t, err)
	assert.Equal(t, "CHMA0000115364", character.Uid)

	os.Remove("data.txt")
}

func TestReadFromFileFailure(t *testing.T) {
	cachedCharacters["uhura"] = Character{
		Uid:     "CHMA0000115364",
		Species: []string{"Human"},
	}
	assert.Nil(t, updateCache())

	character, err := GetCharacterByName("Ashmeet")

	assert.Equal(t, err.Error(), constants.SPECIES_NOT_FOUND)
	assert.Equal(t, "", character.Uid)

	os.Remove("data.txt")
}

func TestReadCharacterWithoutCache(t *testing.T) {
	character, err := GetCharacterByName("uhura")

	assert.Nil(t, err)
	assert.Equal(t, "CHMA0000115364", character.Uid)

	os.Remove("data.txt")
}
