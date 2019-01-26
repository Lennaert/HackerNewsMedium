package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Story struct {
	CreatedAt       time.Time   `json:"created_at"`
	Title           string      `json:"title"`
	URL             string      `json:"url"`
	Author          string      `json:"author"`
	Points          int         `json:"points"`
	StoryText       interface{} `json:"story_text"`
	CommentText     interface{} `json:"comment_text"`
	NumComments     int         `json:"num_comments"`
	StoryID         interface{} `json:"story_id"`
	StoryTitle      interface{} `json:"story_title"`
	StoryURL        interface{} `json:"story_url"`
	ParentID        interface{} `json:"parent_id"`
	CreatedAtI      int         `json:"created_at_i"`
	Tags            []string    `json:"_tags"`
	ObjectID        string      `json:"objectID"`
}

type HnSearchResult struct {
	Hits []Story `json:"hits"`
}

func main() {

	hits := fetchArticlesFromHn()

	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed, color.Bold).SprintFunc()

	for _, story := range hits.Hits {
		fmt.Printf("%s (%d) %s\n<%s>\n\n", blue(story.CreatedAt.String()), story.Points, red(story.Title), story.URL)
	}
}

/**
 * Fetch Articles from the HN search
 * We only want articles from medium.com in the URL.
 *
 */
func fetchArticlesFromHn() HnSearchResult {
	url := "https://hn.algolia.com/api/v1/search_by_date?query=medium.com&restrictSearchableAttributes=url"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Golang-FetchMediumArticles")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	hits := HnSearchResult{}
	jsonErr := json.Unmarshal(body, &hits)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return hits
}


