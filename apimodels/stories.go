package apimodels

type ReaderStarredStoryHashes struct {
	StarredStoryHashes []string `json:"starred_story_hashes"`
	Result             string   `json:"result"`
	Authenticated      bool     `json:"authenticated"`
	UserID             int      `json:"user_id"`
}
