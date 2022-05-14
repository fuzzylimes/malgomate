package malgomate

import (
	"errors"
	"fmt"
	"net/http"
)

// RankingPage is a paginated response page for ranking query results
type RankingPage struct {
	Ranking []Ranking `json:"data"`
	Paging  Paging    `json:"paging"`
}

// ListPage is a paginated response page for a query that returns a list of results
type ListPage struct {
	Listing []Listing `json:"data"`
	Paging  Paging    `json:"paging"`
}

// Anime are the general response objects
type Anime struct {
	ID                     int                `json:"id"`
	Title                  string             `json:"title"`
	MainPicture            *MainPicture       `json:"main_picture,omitempty"`
	AlternativeTitles      *AlternativeTitles `json:"alternative_titles,omitempty"`
	StartDate              string             `json:"start_date,omitempty"`
	EndDate                string             `json:"end_date,omitempty"`
	Synopsis               string             `json:"synopsis,omitempty"`
	Mean                   float64            `json:"mean,omitempty"`
	Rank                   int                `json:"rank,omitempty"`
	Popularity             int                `json:"popularity,omitempty"`
	NumListUsers           int                `json:"num_list_users,omitempty"`
	NumScoringUsers        int                `json:"num_scoring_users,omitempty"`
	Nsfw                   string             `json:"nsfw,omitempty"`
	CreatedAt              string             `json:"created_at,omitempty"`
	UpdatedAt              string             `json:"updated_at,omitempty"`
	MediaType              string             `json:"media_type,omitempty"`
	Status                 string             `json:"status,omitempty"`
	Genres                 []*Genres          `json:"genres,omitempty"`
	NumEpisodes            int                `json:"num_episodes,omitempty"`
	StartSeason            *StartSeason       `json:"start_season,omitempty"`
	Broadcast              *Broadcast         `json:"broadcast,omitempty"`
	Source                 string             `json:"source,omitempty"`
	AverageEpisodeDuration int                `json:"average_episode_duration,omitempty"`
	Rating                 string             `json:"rating,omitempty"`
	Pictures               []*Pictures        `json:"pictures,omitempty"`
	Background             string             `json:"background,omitempty"`
	RelatedAnime           []*RelatedAnime    `json:"related_anime,omitempty"`
	RelatedManga           []interface{}      `json:"related_manga,omitempty"`
	Recommendations        []*Recommendations `json:"recommendations,omitempty"`
	Studios                []*Studios         `json:"studios,omitempty"`
	Statistics             *Statistics        `json:"statistics,omitempty"`
}

// MainPicture contains links to the cover art on MAL
type MainPicture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

// Rank is the current rank of the associated anime based on MAL user scores
type Rank struct {
	Rank int `json:"rank"`
}

// Ranking is a wrapper object for anime objects and their rankings. Similar to Listing,
// just with a Rank property.
type Ranking struct {
	Node Anime `json:"node"`
	Rank Rank  `json:"ranking"`
}

// Listing is a wrapper object for anime objects
type Listing struct {
	Node Anime `json:"node"`
}

// AlternativeTitles are other names that the anime may go by
type AlternativeTitles struct {
	Synonyms []string `json:"synonyms,omitempty"`
	En       string   `json:"en,omitempty"`
	Ja       string   `json:"ja,omitempty"`
}

// Genres are anime genres
type Genres struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// StartSeason is the year and season in which an anime first aired
type StartSeason struct {
	Year   int    `json:"year,omitempty"`
	Season string `json:"season,omitempty"`
}

// Broadcast is when the anime aired
type Broadcast struct {
	DayOfTheWeek string `json:"day_of_the_week,omitempty"`
	StartTime    string `json:"start_time,omitempty"`
}

// Pictures are assocaited images with the anime listing on the MAL website
type Pictures struct {
	Medium string `json:"medium,omitempty"`
	Large  string `json:"large,omitempty"`
}

// Node is a general data object
type Node struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	MainPicture MainPicture `json:"main_picture"`
}

// RelatedAnime are anime that MAL users have said are related to the associated anime
type RelatedAnime struct {
	Node                  Anime  `json:"node,omitempty"`
	RelationType          string `json:"relation_type,omitempty"`
	RelationTypeFormatted string `json:"relation_type_formatted,omitempty"`
}

// Recomendations are anime that MAL users have marked as being similar to the
// associated anime
type Recommendations struct {
	Node               Anime `json:"node,omitempty"`
	NumRecommendations int   `json:"num_recommendations,omitempty"`
}

// Studios are the teams that create the anime
type Studios struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Status are the states an anime can be in on a users list
type Status struct {
	Watching    string `json:"watching,omitempty"`
	Completed   string `json:"completed,omitempty"`
	OnHold      string `json:"on_hold,omitempty"`
	Dropped     string `json:"dropped,omitempty"`
	PlanToWatch string `json:"plan_to_watch,omitempty"`
}

// Statistics are MAL specific datapoints
type Statistics struct {
	Status       Status `json:"status,omitempty"`
	NumListUsers int    `json:"num_list_users,omitempty"`
}

// AnimeQuery is used to perform a general query. Query must be set. Supports fields of the QueryField type
type AnimeQuery struct {
	Query  string
	Limit  int
	Offset int
	Fields QueryFields
}

// DetailsQuery is used to query specific anime by their MAL Ids. The Id must be set. Supports fields of the
// DetailField type
type DetailsQuery struct {
	Id     int
	Fields DetailFields
}

// RankingQuery is used to query for anime rankings. Supports fields of the QueryField type
type RankingQuery struct {
	RankingType RankingType
	Limit       int
	Offset      int
	Fields      QueryFields
}

// SeasonalQuery is used to query for seasonal anime. The Year and the Season must be set. Supports fields of the
// QueryField type
type SeasonalQuery struct {
	Year   int
	Season Season
	Sort   SeasonSort
	Limit  int
	Offset int
	Fields QueryFields
}

// GetDetails retrieves specifics for a given MAL anime Id.
func (c *Client) GetDetails(dq *DetailsQuery) (*Anime, error) {
	// Check for required values
	if dq.Id == 0 {
		return nil, errors.New("missing required parameter: Id must be set")
	}
	// Handle defaults
	if len(dq.Fields) == 0 {
		dq.Fields = BasicDetailQuery
	}

	queryFields := dq.Fields.ToString()
	queryString := fmt.Sprintf("%s/anime/%d?fields=%s", c.BaseURL, dq.Id, queryFields)
	req, err := http.NewRequest(http.MethodGet, queryString, nil)
	if err != nil {
		return nil, err
	}

	res := Anime{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

// GetAnime queries all anime based on a provided string. These queries return a paged list of responses containing
// the fields specified in the intial request object. The query string must be provided. If not included, the
// following default values will be used:
//    * Limit - 100 (max 100)
//    * Offset - 0
//    * Fields - "id,title,main_picture"
func (c *Client) GetAnime(aq *AnimeQuery) (*ListPage, error) {
	// Check for required values
	if aq.Query == "" {
		return nil, errors.New("missing required parameter: Query must be set")
	}
	// Handle defaults
	if aq.Limit == 0 {
		aq.Limit = 100
	} else if aq.Limit > SmallQueryLimit {
		aq.Limit = SmallQueryLimit
	}
	if len(aq.Fields) == 0 {
		aq.Fields = BasicFieldQuery
	}

	queryFields := aq.Fields.ToString()
	queryString := fmt.Sprintf("%s/anime?q=%s&limit=%d&offset=%d&fields=%s", c.BaseURL, aq.Query, aq.Limit, aq.Offset, queryFields)
	req, err := http.NewRequest(http.MethodGet, queryString, nil)
	if err != nil {
		return nil, err
	}

	res := ListPage{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetRanking queries all anime based on rankings on MAL. These queries return a paged list of responses containing
// the fields specified in the intial request object. Rankings are returned based on the provided RankingType value.
// If not included, the following default values will be used:
//    * RankingType - "all"
//    * Limit - 100 (max 500)
//    * Offset - 0
//    * Fields - "id,title,main_picture"
func (c *Client) GetRanking(r *RankingQuery) (*RankingPage, error) {
	// Handle defaults
	if r.RankingType == "" {
		r.RankingType = RankingAll
	}
	if r.Limit == 0 {
		r.Limit = 100
	} else if r.Limit > LargeQueryLimit {
		r.Limit = LargeQueryLimit
	}
	if len(r.Fields) == 0 {
		r.Fields = BasicFieldQuery
	}

	queryFields := r.Fields.ToString()
	queryString := fmt.Sprintf("%s/anime/ranking?ranking_type=%s&limit=%d&offset=%d&fields=%s", c.BaseURL, r.RankingType, r.Limit, r.Offset, queryFields)
	req, err := http.NewRequest(http.MethodGet, queryString, nil)
	if err != nil {
		return nil, err
	}

	res := RankingPage{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetSeason queries for seasonal anime. These queries return a paged list of responses containing the fields specified
// in the initial request object. A year and season MUST be provided. If not included, the following default values
// will be user:
//    * Sort - "anime_score"
//    * Limit - 100 (max 500)
//    * Offset - 0
//    * Fields - "id,title,main_picture"
func (c *Client) GetSeason(q *SeasonalQuery) (*ListPage, error) {
	// Check for required values
	if q.Year == 0 || q.Season == "" {
		return nil, errors.New("missing required parameter: Year and Season must be set")
	}
	// Handle defaults
	if q.Limit == 0 {
		q.Limit = 100
	} else if q.Limit > LargeQueryLimit {
		q.Limit = LargeQueryLimit
	}
	if q.Sort == "" {
		q.Sort = SeasonSortScore
	}
	if len(q.Fields) == 0 {
		q.Fields = BasicFieldQuery
	}

	// Query
	queryFields := q.Fields.ToString()
	queryString := fmt.Sprintf("%s/anime/season/%d/%s?sort=%s&limit=%d&offset=%d&fields=%s", c.BaseURL, q.Year, q.Season, q.Sort, q.Limit, q.Offset, queryFields)
	req, err := http.NewRequest(http.MethodGet, queryString, nil)
	if err != nil {
		return nil, err
	}

	res := ListPage{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetListQS performs a query based on a provided query string. Allows queries to be constructed elsewhere,
// such as the frontend or from a previous/next link. Specifically intended for resouces that return ListPage
// result objects (GetAnime, GetSeason)
func (c *Client) GetListQS(qs string) (*ListPage, error) {
	req, err := http.NewRequest(http.MethodGet, qs, nil)
	if err != nil {
		return nil, err
	}

	res := ListPage{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetRankingQS performs a query based on a provided query string. Allows queries to be constructed elsewhere,
// such as the frontend or from a previous/next link. Specifically intended for Ranking resouce.
func (c *Client) GetRankingQS(qs string) (*RankingPage, error) {
	req, err := http.NewRequest(http.MethodGet, qs, nil)
	if err != nil {
		return nil, err
	}

	res := RankingPage{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
