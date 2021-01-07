package newsblur

type ApiFeed struct {
	ID          int    `json:"id"`
	Ps          int    `json:"ps"`           // positive/focus count
	Nt          int    `json:"nt"`           // neutral/unread count
	Ng          int    `json:"ng"`           // negative/hidden count
	FeedAddress string `json:"feed_address"` // link to the feed (usually .xml/.atom)
	FeedLink    string `json:"feed_link"`    // link to the website
	FeedTitle   string `json:"feed_title"`
}

type ApiStory struct {
	StoryAuthors     string `json:"story_authors"`
	StoryPermalink   string `json:"story_permalink"`
	StoryTimestamp   string `json:"story_timestamp"`
	StoryHash        string `json:"story_hash"`
	ID               string `json:"id"`
	StoryDate        string `json:"story_date"`
	ShortParsedDate  string `json:"short_parsed_date"`
	GUIDHash         string `json:"guid_hash"`
	StoryFeedID      int    `json:"story_feed_id"`
	LongParsedDate   string `json:"long_parsed_date"`
	ReadStatus       int    `json:"read_status"`
	HasModifications bool   `json:"has_modifications"`
	StoryTitle       string `json:"story_title"`
	StoryContent     string `json:"story_content"`
	Starred          bool   `json:"starred"`
}

type LoginInput struct {
	Username string
	Password string
}

type LoginOutput struct {
	Authenticated bool        `json:"authenticated"`
	Code          int         `json:"code"`
	Errors        interface{} `json:"errors"`
	Result        string      `json:"result"`
}

type ReaderFeedsOutput struct {
	Folders []interface{} `json:"folders"`
	Feeds   []ApiFeed     `json:"feeds"`
}

type ReaderFeedInput struct {
	FeedID string
	Page   string
}

type ReaderFeedOutput struct {
	Stories []ApiStory `json:"stories"`
}

type ReaderRiverStoriesInput struct {
	Feeds []string
	Page  string
}

type ReaderRiverStoriesOutput struct {
	Stories []ApiStory `json:"stories"`
}

type MarkStoryHashesAsReadInput struct {
	StoryHash []string
}

type MarkStoryHashesAsReadOutput struct {
}

type MarkStoryHashAsUnreadInput struct {
	StoryHash string
}

type MarkStoryHashAsUnreadOutput struct {
}

type MarkStoryHashAsStarredInput struct {
	StoryHash string
}

type MarkStoryHashAsStarredOutput struct {
}

type MarkStoryHashAsUnstarredInput struct {
	StoryHash string
}

type MarkStoryHashAsUnstarredOutput struct {
}

type ClassifierInput struct {
	FeedID string
}

type ClassifierOutput struct {
	Payload struct {
		Titles map[string]int `json:"titles"`
	} `json:"payload"`
}
