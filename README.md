![malgomate](resources/malgomate.png)
![PkgGoDev](https://pkg.go.dev/badge/github.com/fuzzylimes/malgomate)
![GitHub](https://img.shields.io/github/license/fuzzylimes/malgomate)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/fuzzylimes/malgomate?label=version)

malgomate is a simple go library for the MAL (MyAnimeList) public API. It aims to be a light weight wrapper for making API calls, without a lot of bells and whistles. Grab the data, and do what you want with it.

malgomate does not aim to cover all aspects of the MAL API. It is not interested in user lists, updating of those lists, or any community aspects. It is purely for getting anime data out of MAL. If you are interested in having access to that kind of data, feel free to open an issue and contribute.

## Requirements

In order to use malgomate in your project, you will need to have registered an application in the MAL dev console and aquired an access token. You will need this token in order to make requests.

## Install

Add malgomate to your project by running:
```
go get github.com/fuzzylimes/malgomate@latest
```

## Usage

You can see some basic examples in the `it` folder. Nothing too exciting here:

```go
import mal "github.com/fuzzylimes/malgomate"

func TestListing(t *testing.T) {
	c := mal.NewClient(os.Getenv("MAL_API_KEY"))
	res, err := c.GetAnime(&mal.AnimeQuery{
		Query: "Naruto",
	})
	if err != nil {
		t.Errorf("Unexpected error: %q", err)
	}
}
```

### Helper Types
In order to make it easier to validate incoming requests from the front end, a few helper items exist to validate incoming query data:

| Value                 | Type            | Description                                             |
|-----------------------|-----------------|---------------------------------------------------------|
| SeasonTypeQueries     | SeasonTypes     | List of valid Season values, used in season queries     |
| SeasonSortTypeQueries | SeasonSortTypes | List of valid SeasonSort values, used in season queries |
| RankTypeQueries       | RankingTypes    | List of valid RankType values, used in Ranking queries  |

Each of these values has it's own `.IsValid(str string)` method that can be used to check if an incoming string value is supported for that given query type.


### SubFields
The MAL API provies a way for you specify sub fields for fields that result in an anime response. Currently, this is only supported on a handful of `DetailField` when performing Detail queries using a `DetailsQuery`. The list of supported `DetailField` are as follows:

* RelatedAnime
* Recommendations

In order to add additional fields, you simply chain a `.SubFields()` onto one of the above. It would look something like the following (you can see the full example in the it folder):

```go
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
```

Nothing bad will happen if you run `.SubFields()` on a `DetailField` not in this list, but you also won't get anything else back.

## Support

As mentioned before, malgomate does not support all of the [MAL v2.0 API features](https://myanimelist.net/apiconfig/references/api/v2). Specifically, it only provides interfaces for the following queries:

* [GET Anime List](https://myanimelist.net/apiconfig/references/api/v2#operation/anime_get)
* [GET Anime Details](https://myanimelist.net/apiconfig/references/api/v2#operation/anime_anime_id_get)
* [GET Anime Ranking](https://myanimelist.net/apiconfig/references/api/v2#operation/anime_ranking_get)
* [GET Seasonal Anime](https://myanimelist.net/apiconfig/references/api/v2#operation/anime_season_year_season_get)

### What does malgomate mean?

Nothing - made up word that I thought sounded interesting, that also happened to contain both `MAL` and `GO`.

## TODO List

[X] Handle nested queries for QueryFields<br>
[X] Handle nested queries for DetailFields<br>
[ ] Add UT