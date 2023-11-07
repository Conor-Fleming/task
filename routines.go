package main

import (
	"strings"
	"sync"
)

func fetchNameAndJoke() (string, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var name *NameResponse
	var joke *JokeResponse
	var err error
	errChan := make(chan error)

	//spinning go routines for the requests
	//populating errChan in case of errors
	go func() {
		name, err = getNameData()
		if err != nil {
			errChan <- err
			return
		}
		wg.Done()
	}()
	go func() {
		joke, err = getJokeData()
		if err != nil {
			errChan <- err
			return
		}
		wg.Done()
	}()
	wg.Wait()

	//checking for errors from go routines
	select {
	case e := <-errChan:
		return "", e
	default:
	}

	// Replace name values in Joke with values from Name API
	joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*first", name.FirstName)
	joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*last", name.LastName)
	return joke.Value.Joke, nil
}
