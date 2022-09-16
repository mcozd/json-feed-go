package jsonfeed

import "time"

const CurrentVersion string = "https://jsonfeed.org/version/1.1"

func validStr(str *string) bool {
	return str != nil && len(*str) > 0
}

func hasValidAuthors(authors []Author) bool {
	if authors == nil {
		return true
	}
	for _, author := range authors {
		if author.IsValid() {
			return true
		}
	}
	return false
}

type Feed struct {
	Version     string   `json:"version"`
	Title       string   `json:"title"`
	HomePageUrl *string  `json:"home_page_url"`
	FeedUrl     *string  `json:"feed_url"`
	Description *string  `json:"Description"`
	UserComment *string  `json:"user_comment"`
	NextUrl     *string  `json:"next_url"`
	Icon        *string  `json:"icon"`
	Favicon     *string  `json:"favicon"`
	Language    *string  `json:"language"`
	Expired     *bool    `json:"expired"`
	Authors     []Author `json:"authors,omitempty"`
	Hubs        []Hub    `json:"hubs,omitempty"`
	Items       []Item   `json:"items,omitempty"`
}

func (f Feed) IsEmpty() bool {
	return f.Authors == nil &&
		f.Hubs == nil &&
		f.Items == nil &&
		f.HomePageUrl == nil &&
		f.FeedUrl == nil &&
		f.Description == nil &&
		f.UserComment == nil &&
		f.NextUrl == nil &&
		f.Icon == nil &&
		f.Favicon == nil &&
		f.Language == nil &&
		f.Expired == nil &&
		len(f.Version) == 0 &&
		len(f.Title) == 0
}

func (f Feed) IsValid() bool {
	return validStr(&f.Title) && validStr(&f.Version) &&
		f.Version == CurrentVersion &&
		f.hasValidHubs() &&
		f.hasValidAuthors() &&
		f.hasValidItems()
}

func (f Feed) hasValidAuthors() bool {
	return hasValidAuthors(f.Authors)
}

func (f Feed) hasValidHubs() bool {
	if f.Hubs == nil {
		return true
	}
	for _, hubs := range f.Hubs {
		if hubs.IsValid() {
			return true
		}
	}

	return false
}

func (f Feed) hasValidItems() bool {
	if f.Items == nil {
		return false
	}
	for _, item := range f.Items {
		if item.IsValid() {
			return true
		}
	}

	return false
}

type Author struct {
	Name   *string `json:"name"`
	Url    *string `json:"url"`
	Avatar *string `json:"avatar"`
}

func (a *Author) IsValid() bool {
	return a != nil &&
		(validStr(a.Name) || validStr(a.Url) || validStr(a.Avatar))
}

type Hub struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

func (h *Hub) IsValid() bool {
	return h != nil &&
		validStr(&h.Type) && validStr(&h.Url)
}

type Item struct {
	Id            string       `json:"id"`
	Url           *string      `json:"url"`
	ExternalUrl   *string      `json:"external_url"`
	Title         *string      `json:"title"`
	ContentHtml   *string      `json:"content_html"`
	ContentText   *string      `json:"content_text"`
	Summary       *string      `json:"summary"`
	Image         *string      `json:"image"`
	BannerImage   *string      `json:"banner_image"`
	DatePublished *time.Time   `json:"date_published"`
	DateModified  *time.Time   `json:"date_modified"`
	Language      *string      `json:"language"`
	Authors       []Author     `json:"authors,omitempty"`
	Tags          []string     `json:"tags,omitempty"`
	Attachments   []Attachment `json:"attachments,omitempty"`
}

func (i *Item) isValid() bool {
	return i != nil &&
		validStr(&i.Id) &&
		(validStr(i.ContentHtml) || validStr(i.ContentText))
}

func (i *Item) IsValid() bool {
	return i.isValid() &&
		i.hasValidAuthors() &&
		i.hasValidAttachments()
}

func (i *Item) hasValidAuthors() bool {
	return hasValidAuthors(i.Authors)
}

func (i Item) hasValidAttachments() bool {
	if i.Attachments == nil {
		return true
	}
	for _, attachment := range i.Attachments {
		if attachment.IsValid() {
			return true
		}
	}
	return false
}

type Attachment struct {
	Url               string  `json:"url"`
	MimeType          string  `json:"mime_type"`
	Title             *string `json:"title"`
	SizeInBytes       *int    `json:"size_in_bytes"`
	DurationInSeconds *int    `json:"duration_in_seconds"`
}

func (a *Attachment) IsValid() bool {
	return a != nil &&
		validStr(&a.Url) && validStr(&a.MimeType)
}
