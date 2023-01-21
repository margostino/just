package client

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func Call(url string) ([]byte, error, int) {
	log.Printf("Calling URL %s\n", url)

	response, err := http.Get(url)

	if err != nil {
		log.Printf("error making http request: %s\n", err)
		return nil, errors.New("upstream call failed without status code"), http.StatusInternalServerError
	}

	if response.StatusCode != 200 {
		log.Printf("upstream failed with status code: %d\n", response.StatusCode)
		return nil, errors.New("upstream call failed"), response.StatusCode
	}

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("client: could not read response body: %s\n", err)
		return nil, errors.New("response parse failed"), http.StatusInternalServerError
	}

	return content, nil, http.StatusOK
}
