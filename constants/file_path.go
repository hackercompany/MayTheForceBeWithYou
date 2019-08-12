package constants

import (
	"os"
	"strings"
)

var (

	// DATA_FILE specifies the path of the cache file
	DATA_FILE = "data.txt"

	// KLINGON dictionary data
	KLINGON_DATA = getGenericDir() + "/codex/klingon.json"
)

func getGenericDir() string {
	currDir, _ := os.Getwd()
	currDirSlice := strings.Split(currDir, "/")

	mainDirSlice := make([]string, 0)

	for _, i := range currDirSlice {
		if i != "codex" {
			mainDirSlice = append(mainDirSlice, i)
		}
	}

	return strings.Join(mainDirSlice, "/")
}
