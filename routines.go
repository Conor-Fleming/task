package main

import (
	"errors"
	"strings"
	"sync"
	"time"
)

func fetchNameAndJoke(timeout time.Duration) (string, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var name *NameResponse
	var joke *JokeResponse
	//creating channel for errors
	errChan := make(chan error)

	//spinning go routines for the requests
	//populating errChan in case of errors
	go func() {
		defer wg.Done()
		//implementing a timeout mechanism for requests that take too long
		timer := time.AfterFunc(timeout, func() {
			errChan <- errors.New("name api call timed out")
		})
		nameData, err := getNameData()
		if err != nil {
			errChan <- err
			return
		}
		timer.Stop()
		name = nameData
	}()
	go func() {
		defer wg.Done()
		//implementing a timeout mechanism for requests that take too long
		timer := time.AfterFunc(timeout, func() {
			errChan <- errors.New("joke api timed out")
		})
		jokeData, err := getJokeData()
		if err != nil {
			errChan <- err
			return
		}
		timer.Stop()
		joke = jokeData
	}()

	//waiting for go routines to complete and closing channel
	wg.Wait()
	close(errChan)

	//checking for errors from go routines
	for err := range errChan {
		return "", err
	}

	// ensuring name or joke values are not empty
	if name == nil || joke == nil {
		return "", errors.New("failed to fetch data from apis")
	}

	// Replace name values in Joke with values from Name API
	joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*first", name.FirstName)
	joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*last", name.LastName)

	return joke.Value.Joke, nil
}
