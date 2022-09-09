package jsonfeed_test

import (
	"testing"

	jsonfeed "github.com/mcozd/json-feed-go"
)

var invalidJsons = []string{"", "{}", "[]"}
var validJson = `{
    "version": "https://jsonfeed.org/version/1.1",
    "title": "My Example Feed",
    "home_page_url": "https://example.org/",
    "feed_url": "https://example.org/feed.json",
    "items": [
        {
            "id": "2",
            "content_text": "This is a second item.",
            "url": "https://example.org/second-item"
        },
        {
            "id": "1",
            "content_html": "<p>Hello, world!</p>",
            "url": "https://example.org/initial-post"
        }
    ]
}`

func TestFromString(t *testing.T) {
	t.Parallel()
	t.Run("Should return an error on invalid json", func(t *testing.T) {
		for _, invalidJson := range invalidJsons {
			_, err := jsonfeed.FromString(invalidJson)
			if err == nil {
				t.FailNow()
			}
		}
	})
	t.Run("Should return a filled feed object", func(t *testing.T) {
		feed, err := jsonfeed.FromString(validJson)
		if feed == nil || err != nil {
			t.FailNow()
		}
	})
}
