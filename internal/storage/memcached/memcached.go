package memcached

import (
	"github.com/ehsundar/ghibli_people/pkg/ghp"
)

type basicMemCachedStorage struct {
	otherStorage ghp.PeopleStorage

	cacheInitialized bool
	cache            []*ghp.Person

	cacheByID map[string]*ghp.Person
}

func New(otherStorage ghp.PeopleStorage) ghp.PeopleStorage {
	return &basicMemCachedStorage{
		otherStorage: otherStorage,
		cacheByID:    make(map[string]*ghp.Person),
	}
}

func (s *basicMemCachedStorage) GetAll() ([]*ghp.Person, error) {
	if !s.cacheInitialized {
		people, err := s.otherStorage.GetAll()
		if err != nil {
			return nil, err
		}

		s.cacheInitialized = true
		s.cache = people
	}

	return s.cache, nil
}

func (s *basicMemCachedStorage) Get(ID string) (*ghp.Person, error) {
	person, ok := s.cacheByID[ID]
	if !ok {
		person, err := s.otherStorage.Get(ID)
		if err != nil {
			return nil, err
		}

		s.cacheByID[ID] = person
	}

	return person, nil
}
