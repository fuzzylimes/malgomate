package malgomate

import (
	"fmt"
	"strings"
)

// Season is the name of the season when a show aired
type Season string

// SeasonSort is the value by which to sort the season results
type SeasonSort string

// RankingType is the type by which to retrieve ranking details
type RankingType string

// QueryField are field names to be returned during a query. Used for everything except for detail queries.
type QueryField string

// QueryFields are a collection of QueryField
type QueryFields []QueryField

// ToString converts a slice of QueryField into a comma separate string of fields
func (q QueryFields) ToString() string {
	var sb strings.Builder
	for i, str := range q {
		sb.WriteString(string(str))
		if i < len(q)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

// DetailField are field names to be returned during a details query
type DetailField string

// SubFields attaches a sub query onto a DetailField
func (df DetailField) SubFields(sq *DetailFields) DetailField {
	str := sq.ToString()
	return DetailField(fmt.Sprintf("%s{%s}", df, str))
}

// DetailFields are a collection of DetailField
type DetailFields []DetailField

// ToString converts a slice of DetailField into a comma separate string of fields
func (d DetailFields) ToString() string {
	var sb strings.Builder
	for i, str := range d {
		sb.WriteString(string(str))
		if i < len(d)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

// Common query values
const (
	LargeQueryLimit int = 500
	SmallQueryLimit int = 100
)

// Season values specify the season in seasonal queries
const (
	SeasonWinter Season = "winter"
	SeasonSpring Season = "spring"
	SeasonSummer Season = "summer"
	SeasonFall   Season = "fall"
)

// SeasonTypes are a collection of Season
type SeasonTypes []Season

// IsValid checks to see if the supplied value is a valid Seaon
func (st SeasonTypes) IsValid(str string) bool {
	converted := Season(str)
	for _, v := range st {
		if v == converted {
			return true
		}
	}
	return false
}

// SeasonSort specifies how to sort seasonal queries
const (
	SeasonSortScore SeasonSort = "anime_score"
	SeasonSortUsers SeasonSort = "anime_num_list_users"
)

// SeasonSortTypes are a collection of SeasonSort
type SeasonSortTypes []SeasonSort

// IsValid checks to see if the supplied value is a valid SeasonSortType
func (sst SeasonSortTypes) IsValid(str string) bool {
	converted := SeasonSort(str)
	for _, v := range sst {
		if v == converted {
			return true
		}
	}
	return false
}

// RankingType are the supported ways to query MAL rankings
const (
	RankingAll          RankingType = "all"
	RankingAiring       RankingType = "airing"
	RankingUpcoming     RankingType = "upcoming"
	RankingTv           RankingType = "tv"
	RankingOva          RankingType = "ova"
	RankingMovie        RankingType = "movie"
	RankingSpecial      RankingType = "special"
	RankingByPopularity RankingType = "bypopularity"
	RankingFavorite     RankingType = "favorite"
)

// RankingTypes are a collection of RankingType
type RankingTypes []RankingType

// IsValid checks to see if the supplied value is a valid RankingType
func (rt RankingTypes) IsValid(str string) bool {
	converted := RankingType(str)
	for _, v := range rt {
		if v == converted {
			return true
		}
	}
	return false
}

// QueryField are the supported fields when performing a query
// that results in a List or Ranking type response
const (
	FieldID                     QueryField = "id"
	FieldTitle                  QueryField = "title"
	FieldMainPicture            QueryField = "main_picture"
	FieldAlternativeTitles      QueryField = "alternative_titles"
	FieldStartDate              QueryField = "start_date"
	FieldEndDate                QueryField = "end_date"
	FieldSynopsis               QueryField = "synopsis"
	FieldMean                   QueryField = "mean"
	FieldRank                   QueryField = "rank"
	FieldPopularity             QueryField = "popularity"
	FieldNumListUsers           QueryField = "num_list_users"
	FieldNumScoringUsers        QueryField = "num_scoring_users"
	FieldNsfw                   QueryField = "nsfw"
	FieldGenres                 QueryField = "genres"
	FieldCreatedAt              QueryField = "created_at"
	FieldUpdatedAt              QueryField = "updated_at"
	FieldMediaType              QueryField = "media_type"
	FieldStatus                 QueryField = "status"
	FieldNumEpisodes            QueryField = "num_episodes"
	FieldStartSeason            QueryField = "start_season"
	FieldBroadcast              QueryField = "broadcast"
	FieldSource                 QueryField = "source"
	FieldAverageEpisodeDuration QueryField = "average_episode_duration"
	FieldStudios                QueryField = "studios"
)

// DetailField are the supported fields when performing a query that
// results in a Detail type response
const (
	DetailID                     DetailField = "id"
	DetailTitle                  DetailField = "title"
	DetailMainPicture            DetailField = "main_picture"
	DetailAlternativeTitles      DetailField = "alternative_titles"
	DetailStartDate              DetailField = "start_date"
	DetailEndDate                DetailField = "end_date"
	DetailSynopsis               DetailField = "synopsis"
	DetailMean                   DetailField = "mean"
	DetailRank                   DetailField = "rank"
	DetailPopularity             DetailField = "popularity"
	DetailNumListUsers           DetailField = "num_list_users"
	DetailNumScoringUsers        DetailField = "num_scoring_users"
	DetailNsfw                   DetailField = "nsfw"
	DetailCreatedAt              DetailField = "created_at"
	DetailUpdatedAt              DetailField = "updated_at"
	DetailMediaType              DetailField = "media_type"
	DetailStatus                 DetailField = "status"
	DetailGenres                 DetailField = "genres"
	DetailNumEpisodes            DetailField = "num_episodes"
	DetailStartSeason            DetailField = "start_season"
	DetailBroadcast              DetailField = "broadcast"
	DetailSource                 DetailField = "source"
	DetailAverageEpisodeDuration DetailField = "average_episode_duration"
	DetailRating                 DetailField = "rating"
	DetailPictures               DetailField = "pictures"
	DetailBackground             DetailField = "background"
	DetailRelatedAnime           DetailField = "related_anime"
	DetailRelatedManga           DetailField = "related_manga"
	DetailRecommendations        DetailField = "recommendations"
	DetailStudios                DetailField = "studios"
	DetailStatistics             DetailField = "statistics"
)

// Common QueryFields when running general queries
var (
	BasicFieldQuery QueryFields = []QueryField{
		FieldID,
		FieldTitle,
		FieldMainPicture,
	}
	BasicInfoFieldQuery QueryFields = []QueryField{
		FieldID,
		FieldTitle,
		FieldMainPicture,
		FieldSynopsis,
		FieldStartDate,
		FieldEndDate,
		FieldMean,
	}
)

// Common DetailFields when running detail queries
var (
	BasicDetailQuery DetailFields = []DetailField{
		DetailID,
		DetailTitle,
		DetailMainPicture,
	}
	BasicInfoDetailQuery DetailFields = []DetailField{
		DetailID,
		DetailTitle,
		DetailMainPicture,
		DetailSynopsis,
		DetailStartDate,
		DetailEndDate,
		DetailMean,
	}
)

// Grouping values for easier error checking
var (
	// SeasonTypeQueries are the supported query values used for seasons
	SeasonTypeQueries SeasonTypes = []Season{
		SeasonWinter,
		SeasonSpring,
		SeasonSummer,
		SeasonFall,
	}

	// SeasonSortTypeQueries are the supported query values used when sorting season queries
	SeasonSortTypeQueries SeasonSortTypes = []SeasonSort{
		SeasonSortScore,
		SeasonSortUsers,
	}

	// RankingTypeQueries are the supported query values you can use when querying for rankings
	RankTypeQueries RankingTypes = []RankingType{
		RankingAll,
		RankingAiring,
		RankingUpcoming,
		RankingTv,
		RankingOva,
		RankingMovie,
		RankingSpecial,
		RankingByPopularity,
		RankingFavorite,
	}
)
