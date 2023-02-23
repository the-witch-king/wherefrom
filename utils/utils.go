package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GenerateHttpError(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errorBody, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			errorBody = []byte("Unable to parse response body")
		}

		return fmt.Errorf("Server error.\nStatus Code: %d\nResponse: %v", resp.StatusCode, string(errorBody))
	}

	return nil
}

func IsHttpError(resp *http.Response) bool {
	return resp.StatusCode < 200 || resp.StatusCode > 299
}
