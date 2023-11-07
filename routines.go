package main

import (
	"errors"
	"strings"
	"sync"
)

func fetchNameAndJoke() (string, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var name *NameResponse
	var joke *JokeResponse
	//creating channel for errors with buffer of 2
	errChan := make(chan error, 2)

	//spinning go routines for the requests
	//populating errChan in case of errors
	go func() {
		defer wg.Done()
		nameData, err := getNameData()
		if err != nil {
			errChan <- err
			return
		}

		name = nameData
	}()

	go func() {
		defer wg.Done()
		jokeData, err := getJokeData()
		if err != nil {
			errChan <- err
			return
		}

		joke = jokeData
	}()
	//waiting for go routines to complete and closing channel
	wg.Wait()
	close(errChan)

	//checking for errors from go routines
	for err := range errChan {
		if err != nil {
			return "", err
		}
	}

	// Replace name values in Joke with values from Name API
	if name == nil || joke == nil {
		return "", errors.New("failed to fetch data from apis")
	}

	joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*first", name.FirstName)
	joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*last", name.LastName)
	return joke.Value.Joke, nil
}
