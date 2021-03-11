package filereader

import (
	"encoding/json"
	"github.com/ehsundar/ghibli_people/pkg/ghp"
	"io/ioutil"
)

type basicFileReaderStorage struct {
	fileName string
}

func New(fileName string) ghp.PeopleStorage {
	return &basicFileReaderStorage{
		fileName: fileName,
	}
}

func (s *basicFileReaderStorage) GetAll() ([]*ghp.Person, error) {
	return s.loadFileContents()
}

func (s *basicFileReaderStorage) Get(ID string) (*ghp.Person, error) {
	people, err := s.loadFileContents()
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

func (s *basicFileReaderStorage) loadFileContents() ([]*ghp.Person, error) {
	fileContents, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return nil, err
	}

	var people []*ghp.Person
	err = json.Unmarshal(fileContents, &people)
	if err != nil {
		return nil, err
	}

	return people, nil
}
