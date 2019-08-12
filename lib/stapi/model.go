package stapi

// stCharacter carries entire information about the character
type stCharacter struct {
	Uid     string            `json:"uid"`
	Name    string            `json:"name"`
	Species []speciesResponse `json:"characterSpecies"`
}

// Species for a given Uid are represented
type speciesResponse struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}
