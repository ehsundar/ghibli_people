package apifetcher

import (
	"github.com/ehsundar/ghibli_people/pkg/ghp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicAPIFetcherStorage_GetAllShouldReturnAtLeastOneElement(t *testing.T) {
	s := assert.New(t)

	storage := New("https://ghibliapi.herokuapp.com")
	people, err := storage.GetAll()

	s.Nil(err)
	s.Greater(len(people), 10)
}

func TestBasicAPIFetcherStorage_GetYakulMustSucceed(t *testing.T) {
	s := assert.New(t)

	storage := New("https://ghibliapi.herokuapp.com")
	yakul, err := storage.Get("030555b3-4c92-4fce-93fb-e70c3ae3df8b")

	s.Nil(err)
	s.Equal("Yakul", yakul.Name)
}

func TestBasicAPIFetcherStorage_GetShouldReturnErrorOnInvalidID(t *testing.T) {
	s := assert.New(t)

	storage := New("https://ghibliapi.herokuapp.com")
	person, err := storage.Get("salam!")

	s.NotNil(err)
	s.Nil(person)
	s.Equal(ghp.ErrPersonNotFound, err)
}
