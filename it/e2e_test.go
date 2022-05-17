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

	if len(res.Ranking) < 1 {
		t.Errorf("Expected non zero number of rankings")
	}

	o, _ := json.Marshal(res)
	fmt.Println(string(o))
}

func TestListing(t *testing.T) {
	c := mal.NewClient(os.Getenv("MAL_API_KEY"))
	res, err := c.GetAnime(&mal.AnimeQuery{
		Query: "Naruto",
	})
	if err != nil {
		t.Errorf("Unexpected error: %q", err)
	}

	if len(res.Listing) < 1 {
		t.Errorf("Expected non zero number of rankings")
	}

	o, _ := json.Marshal(res)
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

func TestDetails(t *testing.T) {
	queryId := 10379
	c := mal.NewClient(os.Getenv("MAL_API_KEY"))
	res, err := c.GetDetails(&mal.DetailsQuery{
		Id: queryId,
		Fields: []mal.DetailField{
			mal.DetailRelatedAnime,
			mal.DetailRating,
		},
	})
	if err != nil {
		t.Errorf("Unexpected error: %q", err)
	}

	if res.ID != queryId {
		t.Errorf("Expected non zero number of rankings")
	}

	o, _ := json.Marshal(res)
	fmt.Println(string(o))
}

func TestSubFields(t *testing.T) {
	queryId := 10379
	c := mal.NewClient(os.Getenv("MAL_API_KEY"))
	res, err := c.GetDetails(&mal.DetailsQuery{
		Id: queryId,
		Fields: []mal.DetailField{
			mal.DetailRelatedAnime.SubFields(&mal.DetailFields{
				mal.DetailRank,
			}),
			mal.DetailRating,
			mal.DetailRecommendations.SubFields(&mal.DetailFields{
				mal.DetailRank,
				mal.DetailEndDate,
			}),
		},
	})
	if err != nil {
		t.Errorf("Unexpected error: %q", err)
	}

	for _, v := range res.RelatedAnime {
		if v.Node.Rank == 0 {
			t.Errorf("Unexpected result, should have been set")
		}
	}

	for _, v := range res.Recommendations {
		if v.Node.EndDate == "" {
			t.Errorf("Should have returned end Date")
		}
	}

	o, _ := json.Marshal(res)
	fmt.Println(string(o))
}
