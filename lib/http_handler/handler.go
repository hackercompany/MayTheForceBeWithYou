package http_handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/hackercompany/StarTrek/constants"
	"github.com/hackercompany/StarTrek/logger"
)

// Global http client to handle add requests
var client = http.Client{
	Timeout: time.Duration(10 * time.Second),
}

// RestCall takes care of all http calls for stapi library
func CallRest(url, method, payload string) ([]byte, error) {

	body := strings.NewReader(payload)

	req, _ := http.NewRequest(method, url, body)

	resp, err := client.Do(req)

	// Handling network errors
	if err != nil {

		logger.Print("CallRest", err.Error())
		return nil, err

	}

	// Body is closed to avoid memory leak
	defer resp.Body.Close()

	// Anything other than 200 will raise error
	if resp.StatusCode != 200 {

		logger.Print("CallRest", constants.API_STATUS_INVALID)
		return nil, errors.New(constants.API_STATUS_INVALID)

	}

	content, err := ioutil.ReadAll(resp.Body)

	return content, err

}
