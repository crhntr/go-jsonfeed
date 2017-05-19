package jsonfeed

import "time"

type Feed struct {
	Version     string `json:"version"`
	Title       string `json:"title"`
	HomeURL     string `json:"home_page_url"`
	FeedURL     string `json:"feed_url"`
	Description string `json:"description"`
	UserComment string `json:"user_comment"`
	NextURL     string `json:"next_url"`
	Icon        string `json:"icon"`
	Favicon     string `json:"favicon"`
	Expired     bool   `json:"expired"`
	Items       []Item `json:"items"`
}

type Author struct {
	Name    string `json:"name"`
	SiteURL string `json:"url"`
	Avitar  string `json:"avitar"`
}

type Hub struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Item struct {
	ID             string
	URL            string
	ExternalURL    string
	Title          string
	ContentHTML    string
	Summary        string
	ImageURL       string
	BannerImageURL string
	DatePublished  time.Time
	DateModified   time.Time
	Author         Author
	Tags           []string
	Attachments    []Attachment
}

type Attachment struct {
	URL         string
	MIMEType    string
	Title       string
	SizeInBytes int
	Duration    time.Duration
}
