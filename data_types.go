package main

// NameResponse will contain the first and last name value
type NameResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// JokeResponse will contain the joke value
type JokeResponse struct {
	Value struct {
		Joke string `json:"joke"`
	} `json:"value"`
}
