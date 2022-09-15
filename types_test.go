package jsonfeed_test

import (
	"testing"

	jsonfeed "github.com/mcozd/json-feed-go"
)

var validAuthors = []jsonfeed.Author{
	{Name: jsonfeed.Ptr("Name")},
	{Url: jsonfeed.Ptr("https://")},
	{Avatar: jsonfeed.Ptr("https://www.gravatar")},
	{
		Name: jsonfeed.Ptr("Name"),
		Url:  jsonfeed.Ptr("https://"),
	},
	{
		Name:   jsonfeed.Ptr("Name"),
		Url:    jsonfeed.Ptr("https://"),
		Avatar: jsonfeed.Ptr("https://www.gravatar"),
	},
}

var invalidAuthors = []jsonfeed.Author{
	{},
	{Name: jsonfeed.Ptr("")},
	{Url: jsonfeed.Ptr("")},
	{Avatar: jsonfeed.Ptr("")},
	{
		Name: jsonfeed.Ptr(""),
		Url:  jsonfeed.Ptr(""),
	},
	{
		Name:   jsonfeed.Ptr(""),
		Url:    jsonfeed.Ptr(""),
		Avatar: jsonfeed.Ptr(""),
	},
}

func TestValidAuthor(t *testing.T) {
	t.Parallel()
	t.Run("Should return false for invalid authors", func(t *testing.T) {
		for _, author := range invalidAuthors {
			if author.IsValid() {
				t.FailNow()
			}
		}
	})
	t.Run("Should return true for valid authors", func(t *testing.T) {
		for _, author := range validAuthors {
			if !author.IsValid() {
				t.FailNow()
			}
		}
	})
}

var invalidHubs = []jsonfeed.Hub{
	{},
	{Type: "websub"},
	{Url: "https://"},
}

var validHub = jsonfeed.Hub{Type: "websub", Url: "https://"}

func TestValidHub(t *testing.T) {
	t.Parallel()
	t.Run("Should return false for invalid hubs", func(t *testing.T) {
		for _, hub := range invalidHubs {
			if hub.IsValid() {
				t.FailNow()
			}
		}
	})
	t.Run("Should return true for a valid hub", func(t *testing.T) {
		if !validHub.IsValid() {
			t.FailNow()
		}
	})
}

var invalidAttachments = []jsonfeed.Attachment{
	{},
	{Url: "https://"},
	{MimeType: "image/png"},
}

var validAttachment = jsonfeed.Attachment{Url: "https://", MimeType: "image/png"}

func TestValidAttachment(t *testing.T) {
	t.Parallel()
	t.Run("Should return false for invalid attachments", func(t *testing.T) {
		for _, attachment := range invalidAttachments {
			if attachment.IsValid() {
				t.FailNow()
			}
		}
	})
	t.Run("Should return true for a valid attachment", func(t *testing.T) {
		if !validAttachment.IsValid() {
			t.FailNow()
		}
	})
}

var validItems = []jsonfeed.Item{
	{Id: "1", ContentHtml: jsonfeed.Ptr("<p>content</p><p>more content</p>"), ContentText: jsonfeed.Ptr("Text without html formating.\n But with other formatting.")},
	{Id: "2", ContentHtml: jsonfeed.Ptr("<p>content</p><p>more content</p>")},
	{Id: "3", ContentText: jsonfeed.Ptr(`{ "data": "some data", "meta": "structured data can also be included in this field" }`)},
	{Id: "4", ContentText: jsonfeed.Ptr("content"), Attachments: nil},
	{Id: "5", ContentText: jsonfeed.Ptr("content"), Authors: nil},
}

var invalidItems = []jsonfeed.Item{
	{},
	{Id: "1"},
	{Id: "2", ContentHtml: jsonfeed.Ptr("")},
	{Id: "3", ContentText: jsonfeed.Ptr("")},
	{Id: "4", ContentHtml: jsonfeed.Ptr(""), ContentText: jsonfeed.Ptr("")},
	{Id: "", ContentHtml: jsonfeed.Ptr("<span>content</span>")},
	{Id: "", ContentText: jsonfeed.Ptr("content")},
	{Id: "", ContentHtml: jsonfeed.Ptr("<span>content</span>"), ContentText: jsonfeed.Ptr("content")},
}

var itemsWithInvalidAuthors = []jsonfeed.Item{
	{Id: "1", ContentText: jsonfeed.Ptr("content"), Authors: []jsonfeed.Author{}},
	{Id: "2", ContentText: jsonfeed.Ptr("content"), Authors: []jsonfeed.Author{{}}},
}

var itemsWithInvalidAttachments = []jsonfeed.Item{
	{Id: "1", ContentText: jsonfeed.Ptr("content"), Attachments: []jsonfeed.Attachment{}},
	{Id: "2", ContentText: jsonfeed.Ptr("content"), Attachments: []jsonfeed.Attachment{{}}},
	{Id: "3", ContentText: jsonfeed.Ptr("content"), Attachments: []jsonfeed.Attachment{{Url: "https://"}}},
}

func TestValidItem(t *testing.T) {
	t.Parallel()
	t.Run("Should return false for invalid items", func(t *testing.T) {
		for _, item := range invalidItems {
			if item.IsValid() {
				t.FailNow()
			}
		}
	})
	t.Run("Should return false for items with invalid authors", func(t *testing.T) {
		for _, item := range itemsWithInvalidAuthors {
			if item.IsValid() {
				t.FailNow()
			}
		}
	})
	t.Run("Should return false for items with invalid attachments", func(t *testing.T) {
		for _, item := range itemsWithInvalidAttachments {
			if item.IsValid() {
				t.FailNow()
			}
		}
	})
	t.Run("Should return true for valid items", func(t *testing.T) {
		for _, item := range validItems {
			if !item.IsValid() {
				t.FailNow()
			}
		}
	})
}

var invalidFeeds = []jsonfeed.Feed{
	{Version: "https://jsonfeed.org/version/1.1", Title: "", Items: validItems},
	{Version: "", Title: "Test Feed", Items: validItems},
	{Version: "", Title: "", Items: validItems},
	{Version: "https://jsonfeed.org/version/1.1", Title: "Test Feed", Items: nil},
	{Version: "https://jsonfeed.org/version/1.1", Title: "Test Feed", Items: []jsonfeed.Item{}},
	{Version: "https://jsonfeed.org/version/1.1", Title: "Test Feed", Items: invalidItems},
}

var validFeed = jsonfeed.Feed{Version: "https://jsonfeed.org/version/1.1", Title: "Test Feed", Items: validItems}

func TestValidFeed(t *testing.T) {
	t.Parallel()
	t.Run("Should return false for invalid feed", func(t *testing.T) {
		for _, item := range invalidFeeds {
			if item.IsValid() {
				t.FailNow()
			}
		}
	})
	t.Run("Should return true for a valid feed", func(t *testing.T) {
		if !validFeed.IsValid() {
			t.FailNow()
		}
	})
}
