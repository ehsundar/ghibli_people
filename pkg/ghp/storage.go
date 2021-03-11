package ghp

import "errors"

var (
	ErrPersonNotFound = errors.New("no such person")
)

// Person is a data container for https://ghibliapi.herokuapp.com/#tag/People
type Person struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Age       string `json:"age"`
	EyeColor  string `json:"eye_color"`
	HairColor string `json:"hair_color"`
}

type PeopleStorage interface {
	GetAll() ([]*Person, error)
	Get(ID string) (*Person, error)
}
