package omdb

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func TestSearchByTitle(t *testing.T) {
	Client = &MockClient{}

	apiResp := `{
		"Search": [
		  {
			"Title": "The Dark Knight",
			"Year": "2008",
			"imdbID": "tt0468569",
			"Type": "movie",
			"Poster": "https://m.media-amazon.com/images/M/MV5BMTMxNTMwODM0NF5BMl5BanBnXkFtZTcwODAyMTk2Mw@@._V1_SX300.jpg"
		  },
		  {
			"Title": "The Dark Knight Rises",
			"Year": "2012",
			"imdbID": "tt1345836",
			"Type": "movie",
			"Poster": "https://m.media-amazon.com/images/M/MV5BMTk4ODQzNDY3Ml5BMl5BanBnXkFtZTcwODA0NTM4Nw@@._V1_SX300.jpg"
		  },
		  {
			"Title": "Thor: The Dark World",
			"Year": "2013",
			"imdbID": "tt1981115",
			"Type": "movie",
			"Poster": "https://m.media-amazon.com/images/M/MV5BMTQyNzAwOTUxOF5BMl5BanBnXkFtZTcwMTE0OTc5OQ@@._V1_SX300.jpg"
		  },
		  {
			"Title": "Transformers: Dark of the Moon",
			"Year": "2011",
			"imdbID": "tt1399103",
			"Type": "movie",
			"Poster": "https://m.media-amazon.com/images/M/MV5BMTkwOTY0MTc1NV5BMl5BanBnXkFtZTcwMDQwNjA2NQ@@._V1_SX300.jpg"
		  },
		  {
			"Title": "Dark",
			"Year": "2017–2020",
			"imdbID": "tt5753856",
			"Type": "series",
			"Poster": "https://m.media-amazon.com/images/M/MV5BOTk2NzUyOTctZDdlMS00MDJlLTgzNTEtNzQzYjFhNjA0YjBjXkEyXkFqcGdeQXVyMjg1NDcxNDE@._V1_SX300.jpg"
		  },
		  {
			"Title": "Zero Dark Thirty",
			"Year": "2012",
			"imdbID": "tt1790885",
			"Type": "movie",
			"Poster": "https://m.media-amazon.com/images/M/MV5BMTQ4OTUyNzcwN15BMl5BanBnXkFtZTcwMTQ1NDE3OA@@._V1_SX300.jpg"
		  },
		  {
			"Title": "Dark Shadows",
			"Year": "2012",
			"imdbID": "tt1077368",
			"Type": "movie",
			"Poster": "https://m.media-amazon.com/images/M/MV5BMjc0NzAyMzI1MF5BMl5BanBnXkFtZTcwMTE0NDQ1Nw@@._V1_SX300.jpg"
		  },
		  {
			"Title": "Dark City",
			"Year": "1998",
			"imdbID": "tt0118929",
			"Type": "movie",
			"Poster": "https://m.media-amazon.com/images/M/MV5BMGExOGExM2UtNWM5ZS00OWEzLTllNzYtM2NlMTJlYjBlZTJkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
		  },
		  {
			"Title": "Dark Phoenix",
			"Year": "2019",
			"imdbID": "tt6565702",
			"Type": "movie",
			"Poster": "https://m.media-amazon.com/images/M/MV5BMmZmYTgwZGItNDIxMS00MmRkLWEzODQtYTllNzM0ZWE1NmQ5XkEyXkFqcGdeQXVyODQzNTE3ODc@._V1_SX300.jpg"
		  },
		  {
			"Title": "Terminator: Dark Fate",
			"Year": "2019",
			"imdbID": "tt6450804",
			"Type": "movie",
			"Poster": "https://m.media-amazon.com/images/M/MV5BOWExYzVlZDgtY2E1ZS00NTFjLWFmZWItZjI2NWY5ZWJiNTE4XkEyXkFqcGdeQXVyMTA3MTA4Mzgw._V1_SX300.jpg"
		  }
		],
		"totalResults": "3555",
		"Response": "True"
	  }`

	r := ioutil.NopCloser(bytes.NewReader([]byte(apiResp)))
	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	results, err := SearchByTitle("Dark", "", 1)
	if err != nil {
		t.Fatal(err)
	}
	
	if results.TotalResults != 3555 {
		t.Fatalf("expected 3555 total results, got %v", results.TotalResults)
	}

	if results.Search[4].YearStart != 2017 || results.Search[4].YearEnd != 2020 {
		t.Fatal("error parsing dates")
	}
}
