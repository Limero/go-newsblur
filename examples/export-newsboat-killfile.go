/*
	Example project:
		Export NewsBlur "Intelligence Trainer" data to Newsboat killfile format
*/

package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"

	"github.com/limero/go-newsblur"
)

func addToFilter(filterString string, tag string, classifiers map[string]int) string {
	if len(classifiers) == 0 {
		return filterString
	}

	for classifier, score := range classifiers {
		if score == -1 {
			if filterString != "" {
				filterString += " or "
			}
			filterString += fmt.Sprintf("%s \\\"%s\\\"", tag, classifier)
		}
	}

	return filterString
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s 'username' 'password'", os.Args[0])
		return
	}

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	nb := newsblur.New(&http.Client{
		Jar: cookieJar,
	})

	_, err = nb.Login(os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}

	feeds, err := nb.ReaderFeeds()
	if err != nil {
		panic(err)
	}

	for _, feed := range feeds.Feeds {
		classifiers, err := nb.Classifier(strconv.Itoa(feed.ID))
		if err != nil {
			panic(err)
		}
		filterString := addToFilter("", "author =", classifiers.Payload.Authors)
		filterString = addToFilter(filterString, "title =~", classifiers.Payload.Titles)
		if filterString != "" {
			fmt.Printf("ignore-article \"%s\" \"%s\"\n", feed.FeedAddress, filterString)
		}
	}
}
