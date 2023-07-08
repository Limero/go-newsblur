package newsblur

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Newsblur struct {
	Hostname string
	client   *http.Client
}

func New(client *http.Client) *Newsblur {
	return &Newsblur{
		Hostname: "https://www.newsblur.com",
		client:   client,
	}
}

// Login as an existing user.
// POST /api/login
// https://www.newsblur.com/api#/api/login
func (nb *Newsblur) Login(username, password string) (output *LoginOutput, err error) {
	formData := url.Values{
		"username": {username},
		"password": {password},
	}

	body, err := PostWithBody(nb.client, nb.Hostname+"/api/login", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	if !output.Authenticated {
		return nil, fmt.Errorf("Failed to login to NewsBlur. %v", output.Errors)
	}

	return output, nil
}

// Retrieve a list of feeds to which a user is actively subscribed.
// GET /reader/feeds
// https://www.newsblur.com/api#/reader/feeds
func (nb *Newsblur) ReaderFeeds() (output *ReaderFeedsOutput, err error) {
	body, err := GetWithBody(nb.client, nb.Hostname+"/reader/feeds?v=2")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

// Retrieve stories from a single feed.
// GET /reader/feed/:id
// https://www.newsblur.com/api#/reader/feed/:id
func (nb *Newsblur) ReaderFeed(feedID string, page int) (output *ReaderFeedOutput, err error) {
	if page == 0 {
		page = 1
	}
	body, err := GetWithBody(
		nb.client,
		fmt.Sprintf("%s/reader/feed/%s?page=%d", nb.Hostname, feedID, page),
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

// Retrieve stories from a collection of feeds
// GET /reader/river_stories
// https://www.newsblur.com/api#/reader/river_stories
func (nb *Newsblur) ReaderRiverStories(feeds []string, page int) (output *ReaderRiverStoriesOutput, err error) {
	if page == 0 {
		page = 1
	}

	formData := url.Values{
		"feeds": feeds,
		"page":  {strconv.Itoa(page)},
	}

	body, err := PostWithBody(nb.client, nb.Hostname+"/reader/river_stories", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

// Mark stories as read using their unique story_hash.
// POST /reader/mark_story_hashes_as_read
// https://www.newsblur.com/api#/reader/mark_story_hashes_as_read
func (nb *Newsblur) MarkStoryHashesAsRead(storyHash []string) (output *MarkStoryHashesAsReadOutput, err error) {
	formData := url.Values{
		"story_hash": storyHash,
	}

	body, err := PostWithBody(nb.client, nb.Hostname+"/reader/mark_story_hashes_as_read", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

// Mark a single story as unread using its unique story_hash.
// POST /reader/mark_story_hash_as_unread
// https://www.newsblur.com/api#/reader/mark_story_hash_as_unread
func (nb *Newsblur) MarkStoryHashAsUnread(storyHash string) (output *MarkStoryHashAsUnreadOutput, err error) {
	formData := url.Values{
		"story_hash": {storyHash},
	}

	body, err := PostWithBody(nb.client, nb.Hostname+"/reader/mark_story_hash_as_unread", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

// Mark a story as starred (saved).
// POST /reader/mark_story_hash_as_starred
// https://www.newsblur.com/api#/reader/mark_story_hash_as_starred
func (nb *Newsblur) MarkStoryHashAsStarred(storyHash string) (output *MarkStoryHashAsStarredOutput, err error) {
	formData := url.Values{
		"story_hash": {storyHash},
	}

	body, err := PostWithBody(nb.client, nb.Hostname+"/reader/mark_story_hash_as_starred", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

// Mark a story as unstarred (unsaved).
// POST /reader/mark_story_hash_as_unstarred
// https://www.newsblur.com/api#/reader/mark_story_hash_as_unstarred
func (nb *Newsblur) MarkStoryHashAsUnstarred(storyHash string) (output *MarkStoryHashAsUnstarredOutput, err error) {
	formData := url.Values{
		"story_hash": {storyHash},
	}

	body, err := PostWithBody(nb.client, nb.Hostname+"/reader/mark_story_hash_as_unstarred", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

// Get the intelligence classifiers for a user's site.
// GET /classifier/:id
// https://www.newsblur.com/api#/classifier/:id
func (nb *Newsblur) Classifier(feedID string) (output *ClassifierOutput, err error) {
	body, err := GetWithBody(nb.client, nb.Hostname+"/classifier/"+feedID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}
