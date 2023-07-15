package newsblur

import "encoding/json"

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
