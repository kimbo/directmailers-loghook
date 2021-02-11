package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://print.directmailers.com/api/v1"
)

func New(username string, password string) *API {
	return &API{
		username: username,
		password: password,
		Client:   http.DefaultClient,
	}
}

type API struct {
	Client   *http.Client
	username string
	password string
}

func (a *API) CreateLetter(l LetterRequest) (LetterResponse, error) {
	reqBody, err := marshalNoEscapeHTML(l)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/letter/", baseURL), bytes.NewReader(reqBody))
	if err != nil {
		return LetterResponse{}, err
	}

	req.SetBasicAuth(a.username, a.password)
	req.Header.Set("Content-Type", "application/json")
	res, err := a.Client.Do(req)
	if err != nil {
		return LetterResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		return LetterResponse{}, fmt.Errorf("got unexpected status code %v: %v", res.StatusCode, string(body))
	}

	var lr LetterResponse
	if err := json.NewDecoder(res.Body).Decode(&lr); err != nil {
		return LetterResponse{}, err
	}

	return lr, nil
}

func (a *API) CreatePostcard(p PostcardRequest) (PostcardResponse, error) {
	reqBody, err := marshalNoEscapeHTML(p)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/postcard/", baseURL), bytes.NewReader(reqBody))
	if err != nil {
		return PostcardResponse{}, err
	}
	req.SetBasicAuth(a.username, a.password)
	req.Header.Set("Content-Type", "application/json")
	res, err := a.Client.Do(req)
	if err != nil {
		return PostcardResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		return PostcardResponse{}, fmt.Errorf("got unexpected status code %v: %v", res.StatusCode, string(body))
	}

	var pr PostcardResponse
	if err := json.NewDecoder(res.Body).Decode(&pr); err != nil {
		return PostcardResponse{}, err
	}

	return pr, nil
}

func marshalNoEscapeHTML(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
