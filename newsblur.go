package newsblur

import (
	"encoding/json"
	"net/http"
	"net/url"
)

var server = "https://www.newsblur.com"

/*
	Login as an existing user.
	POST /api/login
	https://www.newsblur.com/api#/api/login
*/
func ApiLogin(client *http.Client, input *LoginInput) (output *LoginOutput, err error) {
	formData := url.Values{
		"username": {input.Username},
		"password": {input.Password},
	}

	body, err := PostWithBody(client, server+"/api/login", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

/*
	Retrieve a list of feeds to which a user is actively subscribed.
	GET /reader/feeds
	https://www.newsblur.com/api#/reader/feeds
*/
func ApiReaderFeeds(client *http.Client) (output *ReaderFeedsOutput, err error) {
	body, err := GetWithBody(client, server+"/reader/feeds?v=2")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

/*
	Retrieve stories from a single feed.
	GET /reader/feed/:id
	https://www.newsblur.com/api#/reader/feed/:id
*/
func ApiReaderFeed(client *http.Client, input *ReaderFeedInput) (output *ReaderFeedOutput, err error) {
	body, err := GetWithBody(client, server+"/reader/feed/"+input.FeedID+"?page="+input.Page)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

/*
	Retrieve stories from a collection of feeds
	GET /reader/river_stories
	https://www.newsblur.com/api#/reader/river_stories
*/
func ApiReaderRiverStories(client *http.Client, input *ReaderRiverStoriesInput) (output *ReaderRiverStoriesOutput, err error) {
	formData := url.Values{
		"feeds": input.Feeds,
		"page":  {input.Page},
	}

	body, err := PostWithBody(client, server+"/reader/river_stories", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

/*
	Mark stories as read using their unique story_hash.
	POST /reader/mark_story_hashes_as_read
	https://www.newsblur.com/api#/reader/mark_story_hashes_as_read
*/
func ApiMarkStoryHashesAsRead(client *http.Client, input *MarkStoryHashesAsReadInput) (output *MarkStoryHashesAsReadOutput, err error) {
	formData := url.Values{
		"story_hash": input.StoryHash,
	}

	body, err := PostWithBody(client, server+"/reader/mark_story_hashes_as_read", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

/*
	Mark a single story as unread using its unique story_hash.
	POST /reader/mark_story_hash_as_unread
	https://www.newsblur.com/api#/reader/mark_story_hash_as_unread
*/
func ApiMarkStoryHashAsUnread(client *http.Client, input *MarkStoryHashAsUnreadInput) (output *MarkStoryHashAsUnreadOutput, err error) {
	formData := url.Values{
		"story_hash": {input.StoryHash},
	}

	body, err := PostWithBody(client, server+"/reader/mark_story_hash_as_unread", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

/*
	Mark a story as starred (saved).
	POST /reader/mark_story_hash_as_starred
	https://www.newsblur.com/api#/reader/mark_story_hash_as_starred
*/
func ApiMarkStoryHashAsStarred(client *http.Client, input *MarkStoryHashAsStarredInput) (output *MarkStoryHashAsStarredOutput, err error) {
	formData := url.Values{
		"story_hash": {input.StoryHash},
	}

	body, err := PostWithBody(client, server+"/reader/mark_story_hash_as_starred", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}

/*
	Mark a story as unstarred (unsaved).
	POST /reader/mark_story_hash_as_unstarred
	https://www.newsblur.com/api#/reader/mark_story_hash_as_unstarred
*/
func ApiMarkStoryHashAsUnstarred(client *http.Client, input *MarkStoryHashAsUnstarredInput) (output *MarkStoryHashAsUnstarredOutput, err error) {
	formData := url.Values{
		"story_hash": {input.StoryHash},
	}

	body, err := PostWithBody(client, server+"/reader/mark_story_hash_as_unstarred", formData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}
