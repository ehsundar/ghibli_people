package memcached

import (
	"github.com/ehsundar/ghibli_people/mocks"
	"github.com/ehsundar/ghibli_people/pkg/ghp"
	"github.com/stretchr/testify/suite"
	"testing"
)

type basicFileReaderStorageTestSuite struct {
	suite.Suite
	person1 *ghp.Person
	person2 *ghp.Person
}

func TestBasicFileReaderStorage(t *testing.T) {
	suite.Run(t, &basicFileReaderStorageTestSuite{})
}

func (s *basicFileReaderStorageTestSuite) SetupSuite() {
	s.person1 = &ghp.Person{
		ID:   "id1",
		Name: "Amir",
	}
	s.person2 = &ghp.Person{
		ID:   "id2",
		Name: "Ehsan",
	}
}

func (s *basicFileReaderStorageTestSuite) TearDownSuite() {
}

func (s *basicFileReaderStorageTestSuite) TestBasicMemCachedStorage_GetAllShouldPassExactThing() {
	mockOtherStorage := &mocks.PeopleStorage{}
	mockOtherStorage.
		On("GetAll").
		Return([]*ghp.Person{s.person1, s.person2}, nil)

	storage := New(mockOtherStorage)
	people, err := storage.GetAll()

	s.Nil(err)
	s.Equal(len(people), 2)
	s.Equal(people[0].ID, "id1")
}

func (s *basicFileReaderStorageTestSuite) TestBasicMemCachedStorage_GetAllShouldUseCacheForRepetitiveCalls() {
	mockOtherStorage := &mocks.PeopleStorage{}
	mockOtherStorage.
		On("GetAll").
		Return([]*ghp.Person{s.person1, s.person2}, nil).
		Once()

	storage := New(mockOtherStorage)

	for i := 0; i < 10; i++ {
		people, err := storage.GetAll()
		s.Nil(err)
		s.Equal(len(people), 2)
		s.Equal(people[0].ID, "id1")
	}

	mockOtherStorage.AssertExpectations(s.T())
}
