package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

// NameData makes request to joke api and unmarshals the json response to our NameResponse struct
func getNameData() (*NameResponse, error) {
	//Calling Api for the random name
	nameResp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		log.Print("Getting name api response: ", err)
		return nil, err
	}
	defer nameResp.Body.Close()

	//reading response to byte slice
	nameRespData, err := io.ReadAll(nameResp.Body)
	if err != nil {
		log.Print("Read Response Data: ", err)
		return nil, err
	}

	// Unmarshalling JSON to name variable of type NameResponse (declared above)
	name := &NameResponse{}
	err = json.Unmarshal(nameRespData, &name)
	if err != nil {
		log.Print("Unmarshal Name: ", err)
		return nil, err
	}

	return name, nil
}

// JokeData makes request to joke api and unmarshals the json response to our JokeResponse struct
func getJokeData() (*JokeResponse, error) {
	//using url package to build string because of
	url := url.URL{
		Scheme:   "http",
		Host:     "joke.loc8u.com:8888",
		Path:     "/joke",
		RawQuery: "limitTo=nerdy&firstName=*first&lastName=*last",
	}

	jokeURL := url.String()

	//Calling Api for the random joke
	jokeResp, err := http.Get(jokeURL)
	if err != nil {
		log.Print("Getting joke api response: ", err)
		return nil, err
	}
	defer jokeResp.Body.Close()

	//reading response to byte slice
	jokeRespData, err := io.ReadAll(jokeResp.Body)
	if err != nil {
		log.Print("Read Response Data: ", err)
		return nil, err
	}

	//Unmarshalling JSON to joke variable of type JokeResponse
	joke := &JokeResponse{}
	err = json.Unmarshal(jokeRespData, &joke)
	if err != nil {
		log.Print("Unmarshal Joke: ", err)
		return nil, err
	}

	return joke, nil
}
