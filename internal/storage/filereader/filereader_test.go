package filereader

import (
	"encoding/json"
	"github.com/ehsundar/ghibli_people/pkg/ghp"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"testing"
)

const testDataFileName = "test_data.json"

type basicFileReaderStorageTestSuite struct {
	suite.Suite
	workingStorage ghp.PeopleStorage
	brokenStorage  ghp.PeopleStorage
}

func TestBasicFileReaderStorage(t *testing.T) {
	suite.Run(t, &basicFileReaderStorageTestSuite{})
}

func (s *basicFileReaderStorageTestSuite) SetupSuite() {
	people := []*ghp.Person{
		{
			ID:   "030555b3-4c92-4fce-93fb-e70c3ae3df8b",
			Name: "Yakul",
		},
		{
			ID:   "test-test",
			Name: "Amir!",
		},
	}
	encodedData, err := json.Marshal(people)
	s.Nil(err)

	err = ioutil.WriteFile(testDataFileName, encodedData, 777)
	s.Nil(err)

	s.workingStorage = New(testDataFileName)
	s.brokenStorage = New("این ره که تو میروی به ترکستان است.json")
}

func (s *basicFileReaderStorageTestSuite) TearDownSuite() {
	err := os.Remove(testDataFileName)
	s.Nil(err)
}

func (s *basicFileReaderStorageTestSuite) TestBasicFileReaderStorage_GetAllShouldReturnAtLeastOneElement() {
	people, err := s.workingStorage.GetAll()

	s.Nil(err)
	s.Greater(len(people), 1)
}

func (s *basicFileReaderStorageTestSuite) TestBasicFileReaderStorage_GetYakulMustSucceed() {
	yakul, err := s.workingStorage.Get("030555b3-4c92-4fce-93fb-e70c3ae3df8b")

	s.Nil(err)
	s.Equal("Yakul", yakul.Name)
}

func (s *basicFileReaderStorageTestSuite) TestBasicFileReaderStorage_GetShouldReturnErrorOnInvalidID() {
	person, err := s.workingStorage.Get("salam!")

	s.NotNil(err)
	s.Nil(person)
	s.Equal(ghp.ErrPersonNotFound, err)
}

func (s *basicFileReaderStorageTestSuite) TestBasicFileReaderStorage_GetAllShouldReturnErrorOnInvalidFileName() {
	person, err := s.brokenStorage.GetAll()

	s.NotNil(err)
	s.Nil(person)
}

func (s *basicFileReaderStorageTestSuite) TestBasicFileReaderStorage_GetShouldReturnErrorOnInvalidFileName() {
	person, err := s.brokenStorage.Get("whatever")

	s.NotNil(err)
	s.Nil(person)
}
