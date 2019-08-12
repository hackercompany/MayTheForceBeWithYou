package codex

import (
	"testing"

	"github.com/hackercompany/MayForceBeWithYou/constants"

	"github.com/stretchr/testify/assert"
)

func TestTranslateToKlingonSuccess(t *testing.T) {
	k := NewKlingon()

	hexValue, err := k.Translate("Ashmeet")
	assert.Nil(t, err)
	assert.Equal(t, hexValue, "0xF8D0 0xF8E2 0xF8D6 0xF8DA 0xF8D4 0xF8D4 0xF8E3")

	hexValue, err = k.Translate("Test")
	assert.Nil(t, err)
	assert.Equal(t, hexValue, "0xF8E3 0xF8D4 0xF8E2 0xF8E3")

	hexValue, err = k.Translate("qQqQ")
	assert.Nil(t, err)
	assert.Equal(t, hexValue, "0xF8DF 0xF8E0 0xF8DF 0xF8E0")
}

func TestTranslateToKligonFailure(t *testing.T) {
	k := NewKlingon()

	hexValue, err := k.Translate("Jexia")
	assert.Equal(t, hexValue, "")
	assert.EqualError(t, err, constants.INVALID_INPUT)

	hexValue, err = k.Translate("asldlksajd")
	assert.Equal(t, hexValue, "")
	assert.EqualError(t, err, constants.INVALID_INPUT)
}

func TestTranslateForEachCharacterSuccess(t *testing.T) {
	k := NewKlingon()

	allCharacters := []string{
		"a", "b", "ch", "d", "e", "gh", "h", "i", "j", "l", "m", "n", "ng", "o", "p", "q", "Q", "r", "s", "t", "tlh", "u", "v", "w", "y", "'", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ",", ".", " ",
	}

	translatedCharacters := []string{
		"0xF8D0", "0xF8D1", "0xF8D2", "0xF8D3", "0xF8D4", "0xF8D5", "0xF8D6", "0xF8D7", "0xF8D8", "0xF8D9", "0xF8DA", "0xF8DB", "0xF8DC", "0xF8DD", "0xF8DE", "0xF8DF", "0xF8E0", "0xF8E1", "0xF8E2", "0xF8E3", "0xF8E4", "0xF8E5", "0xF8E6", "0xF8E7", "0xF8E8", "0xF8E9", "0xF8F0", "0xF8F1", "0xF8F2", "0xF8F3", "0xF8F4", "0xF8F5", "0xF8F6", "0xF8F7", "0xF8F8", "0xF8F9", "0xF8FD", "0xF8FE", "0x0020",
	}

	for index := range allCharacters {
		hexValue, err := k.Translate(allCharacters[index])
		assert.Nil(t, err)
		assert.Equal(t, hexValue, translatedCharacters[index])
	}

}

func TestTranslateForEachCharacterFail(t *testing.T) {
	k := NewKlingon()

	failCases := []string{
		"c", "g", "k", "x", "z",
	}

	for index := range failCases {
		hexValue, err := k.Translate(failCases[index])
		assert.EqualError(t, err, constants.INVALID_INPUT)
		assert.Equal(t, hexValue, "")
	}

}
