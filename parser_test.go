package jsonfeed_test

import (
	"reflect"
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

var validJsonWithFeedAndItemExtensions = `{
    "version": "https://jsonfeed.org/version/1.1",
    "title": "My Example Feed",
    "home_page_url": "https://example.org/",
    "feed_url": "https://example.org/feed.json",
	"_extra1": "test",
	"_extra2": 123,
    "items": [
        {
            "id": "2",
            "content_text": "This is a second item.",
            "url": "https://example.org/second-item",
            "_lead": "Other Person",
            "_stars": 567,
            "_likes": 1234
        },
        {
            "id": "1",
            "content_html": "<p>Hello, world!</p>",
            "url": "https://example.org/initial-post",
            "_lead": "Some Person",
            "_stars": 1
        }
    ]
}`

type FeedExtensions struct {
	Extra1 *string `json:"_extra1"`
	Extra2 *int    `json:"_extra2"`
}

type ExtendedFeed struct {
	jsonfeed.Feed
	FeedExtensions
}

type ItemExtensions struct {
	Lead  *string `json:"_lead"`
	Stars *int    `json:"_stars"`
	Likes *int    `json:"_likes"`
}

type ExtendedItem struct {
	jsonfeed.Item
	ItemExtensions
}

type FeedWithExtendedItems struct {
	jsonfeed.Feed
	Items []ExtendedItem `json:"items,omitempty"`
}

type ExtendedFeedWithExtendedItems struct {
	ExtendedFeed
	Items []ExtendedItem `json:"items,omitempty"`
}

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

func TestUnmarshalFromString(t *testing.T) {
	t.Parallel()
	t.Run("Should return an extended feed object", func(t *testing.T) {
		feed := ExtendedFeed{}
		err := jsonfeed.UnmarshalFromString(validJsonWithFeedAndItemExtensions, &feed)
		if err != nil {
			t.FailNow()
		}
		if *feed.Extra1 != "test" || *feed.Extra2 != 123 {
			t.FailNow()
		}
	})
	t.Run("Should return a feed object with extended items", func(t *testing.T) {
		feed := FeedWithExtendedItems{}
		err := jsonfeed.UnmarshalFromString(validJsonWithFeedAndItemExtensions, &feed)
		if err != nil {
			t.FailNow()
		}

		wantedItem1 := ItemExtensions{Lead: jsonfeed.Ptr("Other Person"), Stars: jsonfeed.Ptr(567), Likes: jsonfeed.Ptr(1234)}
		wantedItem2 := ItemExtensions{Lead: jsonfeed.Ptr("Some Person"), Stars: jsonfeed.Ptr(1), Likes: nil}
		if !reflect.DeepEqual(feed.Items[0].ItemExtensions, wantedItem1) {
			t.FailNow()
		}
		if !reflect.DeepEqual(feed.Items[1].ItemExtensions, wantedItem2) {
			t.FailNow()
		}
	})
	t.Run("Should return a feed object with extensions, containing items with extensions", func(t *testing.T) {
		feed := ExtendedFeedWithExtendedItems{}
		err := jsonfeed.UnmarshalFromString(validJsonWithFeedAndItemExtensions, &feed)
		if err != nil {
			t.FailNow()
		}

		if *feed.Extra1 != "test" || *feed.Extra2 != 123 {
			t.FailNow()
		}

		wantedItem1 := ItemExtensions{Lead: jsonfeed.Ptr("Other Person"), Stars: jsonfeed.Ptr(567), Likes: jsonfeed.Ptr(1234)}
		wantedItem2 := ItemExtensions{Lead: jsonfeed.Ptr("Some Person"), Stars: jsonfeed.Ptr(1), Likes: nil}
		if !reflect.DeepEqual(feed.Items[0].ItemExtensions, wantedItem1) {
			t.FailNow()
		}
		if !reflect.DeepEqual(feed.Items[1].ItemExtensions, wantedItem2) {
			t.FailNow()
		}
	})
}
