package it_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	mal "github.com/fuzzylimes/malgomate"
)

func TestRanking(t *testing.T) {
	c := mal.NewClient(os.Getenv("MAL_API_KEY"))
	fmt.Println(c)
	res, err := c.GetRanking(&mal.RankingQuery{})
	fmt.Println(err)
	if err != nil {
		t.Errorf("Unexpected error: %q", err)
	}

	fmt.Println(res)

	if len(res.Ranking) < 1 {
		t.Errorf("Expected non zero number of rankings")
	}

}

func TestListing(t *testing.T) {
	c := mal.NewClient(os.Getenv("MAL_API_KEY"))
	res, err := c.GetAnime(&mal.AnimeQuery{
		Query: "Naruto",
	})
	if err != nil {
		t.Errorf("Unexpected error: %q", err)
	}
	fmt.Println(res)

	if len(res.Listing) < 1 {
		t.Errorf("Expected non zero number of rankings")
	}

	o, _ := json.Marshal(res)
	fmt.Println(string(o))

	o, _ = json.Marshal(res.Listing[0].Node.MainPicture)
	fmt.Println(string(o))
}

func TestSeasonal(t *testing.T) {
	c := mal.NewClient(os.Getenv("MAL_API_KEY"))
	res, err := c.GetSeason(&mal.SeasonalQuery{
		Year:   2022,
		Season: mal.SeasonWinter,
		Limit:  10,
	})
	if err != nil {
		t.Errorf("Unexpected error: %q", err)
	}

	fmt.Println(res)

	if len(res.Listing) < 1 {
		t.Errorf("Expected non zero number of rankings")
	}

	o, _ := json.Marshal(res)
	fmt.Println(string(o))

	next := mal.ListPage{}
	if n := res.Paging.HasNext(); n == true {
		err = c.GetNextPage(&res.Paging, &next)
	}
	if err != nil {
		t.Errorf("Unexpected error: %q", err)
	}

	o, _ = json.Marshal(next)
	fmt.Println(string(o))
}
