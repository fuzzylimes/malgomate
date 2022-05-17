package malgomate

import (
	"fmt"
	"testing"
)

func TestRankingTypeIsValid(t *testing.T) {
	testCases := []struct {
		in       string
		expected bool
	}{
		{"Bad", false},
		{"all", true},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			if got := RankTypeQueries.IsValid(tc.in); got != tc.expected {
				t.Errorf("Expected %t, got %t", tc.expected, got)
			}
		})
	}
}

func TestSeasonTypeIsValid(t *testing.T) {
	testCases := []struct {
		in       string
		expected bool
	}{
		{"springSeason", false},
		{"spring", true},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			if got := SeasonTypeQueries.IsValid(tc.in); got != tc.expected {
				t.Errorf("Expected %t, got %t", tc.expected, got)
			}
		})
	}
}

func TestSeasonSortTypeIsValid(t *testing.T) {
	testCases := []struct {
		in       string
		expected bool
	}{
		{"animeScore", false},
		{"anime_score", true},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			if got := SeasonSortTypeQueries.IsValid(tc.in); got != tc.expected {
				t.Errorf("Expected %t, got %t", tc.expected, got)
			}
		})
	}
}

func TestDetailSubFields(t *testing.T) {
	testCases := []struct {
		root     DetailField
		in       *DetailFields
		expected string
	}{
		{DetailTitle, &DetailFields{DetailBroadcast}, "title{broadcast}"},
		{DetailBackground, &DetailFields{DetailBroadcast, DetailAlternativeTitles}, "background{broadcast,alternative_titles}"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			if got := tc.root.SubFields(tc.in); got != DetailField(tc.expected) {
				t.Errorf("Expected %s, got %s", tc.expected, got)
			}
		})
	}
}
