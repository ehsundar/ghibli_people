package main

import (
	"fmt"
	"github.com/ehsundar/ghibli_people/internal/storage/apifetcher"
	"github.com/ehsundar/ghibli_people/internal/storage/filereader"
	"github.com/ehsundar/ghibli_people/pkg/ghp"
	"os"
)

func main() {
	mode := os.Args[1]

	var fetcher ghp.PeopleStorage

	switch mode {
	case "api":
		fetcher = apifetcher.New("https://ghibliapi.herokuapp.com")
	case "file":
		fetcher = filereader.New("static_people.json")
	default:
		panic("unknown running mode")
	}

	yakul, err := fetcher.Get("030555b3-4c92-4fce-93fb-e70c3ae3df8b")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", yakul)
}
