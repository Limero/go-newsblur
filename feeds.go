package newsblur

import "encoding/json"

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
