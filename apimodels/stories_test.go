package apimodels

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReaderStarredStoryHashes(t *testing.T) {
	in := `
{
  "starred_story_hashes": [
    "1111111:aaaaaa",
    "2222222:bbbbbb"
  ],
  "result": "ok",
  "authenticated": true,
  "user_id": 500000
}
`

	var output ReaderStarredStoryHashes

	err := json.Unmarshal([]byte(in), &output)
	require.NoError(t, err)

	assert.Len(t, output.StarredStoryHashes, 2)
}
