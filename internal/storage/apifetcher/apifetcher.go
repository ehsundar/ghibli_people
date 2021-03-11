package apifetcher

import (
	"encoding/json"
	"fmt"
	"github.com/ehsundar/ghibli_people/pkg/ghp"
	"io/ioutil"
	"net/http"
)

type basicAPIFetcherStorage struct {
	baseURL string
}

func New(baseURL string) ghp.PeopleStorage {
	return &basicAPIFetcherStorage{
		baseURL: baseURL,
	}
}

func (s *basicAPIFetcherStorage) GetAll() ([]*ghp.Person, error) {
	return s.fetchFromAPI()
}

func (s *basicAPIFetcherStorage) Get(ID string) (*ghp.Person, error) {
	people, err := s.fetchFromAPI()
	if err != nil {
		return nil, err
	}

	for _, p := range people {
		if p.ID == ID {
			return p, nil
		}
	}

	return nil, ghp.ErrPersonNotFound
}

func (s *basicAPIFetcherStorage) fetchFromAPI() ([]*ghp.Person, error) {
	url := fmt.Sprintf("%s/people", s.baseURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(resp.Body)

	var people []*ghp.Person
	err = json.Unmarshal(contents, &people)
	if err != nil {
		return nil, err
	}

	return people, nil
}
