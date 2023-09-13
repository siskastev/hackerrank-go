package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
 * Complete the 'getAuthorHistory' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING author as parameter.
 *
 * Base urls:
 *   https://jsonmock.hackerrank.com/api/article_users?username=
 *   https://jsonmock.hackerrank.com/api/articles?author=
 *
 */

type ArticleUser struct {
	Data []struct {
		ID              int    `json:"id"`
		Username        string `json:"username"`
		About           string `json:"about"`
		SubmissionCount int    `json:"submission_count"`
		CommentCount    int    `json:"comment_count"`
	} `json:"data"`
}

type Article struct {
	Data []struct {
		Title      string `json:"title"`
		ShortTitle string `json:"short_title"`
	} `json:"data"`
	TotalPages int `json:"total_pages"`
}

func getAuthorHistory(author string) []string {
	history := []string{}

	// Make request to article_users API to get user ID
	userURL := fmt.Sprintf("https://jsonmock.hackerrank.com/api/article_users?username=%s", author)
	userResp, err := http.Get(userURL)
	if err != nil {
		log.Fatal(err)
	}
	defer userResp.Body.Close()

	userBody, err := ioutil.ReadAll(userResp.Body)
	if err != nil {
		log.Fatal(err)
	}

	userRespData := new(ArticleUser)
	err = json.Unmarshal(userBody, &userRespData)
	if err != nil {
		panic(err)
	}

	if len(userRespData.Data) == 0 {
		return history
	}

	for _, user := range userRespData.Data {
		if user.About != "" {
			history = append(history, user.About)
		}
	}

	// Make request to articles API to get article history
	articlesURL := fmt.Sprintf("https://jsonmock.hackerrank.com/api/articles?author=%s", author)
	page := 1
	totalPages := 1

	for page <= totalPages {
		articlesResp, err := http.Get(fmt.Sprintf("%s&page=%d", articlesURL, page))
		if err != nil {
			panic(err)
		}
		defer articlesResp.Body.Close()

		articlesBody, err := ioutil.ReadAll(articlesResp.Body)
		if err != nil {
			panic(err)
		}

		articlesRespData := new(Article)

		err = json.Unmarshal(articlesBody, &articlesRespData)
		if err != nil {
			panic(err)
		}

		if len(articlesRespData.Data) == 0 {
			return history
		}

		for _, article := range articlesRespData.Data {
			title := article.Title
			if title == "" {
				sortTitle := article.ShortTitle
				if sortTitle != "" {
					history = append(history, sortTitle)
				}
			} else {
				history = append(history, title)
			}
		}

		totalPages = articlesRespData.TotalPages
		page++
	}

	return history
}

func main() {
	var author string

	fmt.Scanln(&author)

	result := getAuthorHistory(author)

	for _, resultItem := range result {
		fmt.Println(resultItem)
	}
}
