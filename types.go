package jsonfeed

import "time"

type Feed struct {
	Version     string `json:"version"`
	Title       string `json:"title"`
	HomeURL     string `json:"home_page_url,optional"`
	FeedURL     string `json:"feed_url,optional"`
	Description string `json:"description,optional"`
	UserComment string `json:"user_comment,optional"`
	NextURL     string `json:"next_url,optional"`
	Icon        string `json:"icon,optional"`
	Favicon     string `json:"favicon,optional"`
	Expired     bool   `json:"expired,optional"`
	Items       []Item `json:"items,optional"`

	Hubs map[string]interface{} `json:"hubs,optional"`
}

type Author struct {
	Name    string `json:"name,optional"`
	SiteURL string `json:"url,optional"`
	Avitar  string `json:"avitar,optional"`
}

type Hub struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Item struct {
	ID             string       `json:"id"`
	URL            string       `json:"url,optional"`
	ExternalURL    string       `json:"external_url,optional"`
	Title          string       `json:"title,optional"`
	ContentHTML    string       `json:"content_html,optional"`
	ContentText    string       `json:"content_text,optional"`
	Summary        string       `json:"summary,optional"`
	ImageURL       string       `json:"image,optional"`
	BannerImageURL string       `json:"banner_image,optional"`
	DatePublished  time.Time    `json:"date_published,optional"`
	DateModified   time.Time    `json:"date_modified,optional"`
	Author         Author       `json:"author,optional"`
	Tags           []string     `json:"tags,optional"`
	Attachments    []Attachment `json:"attachments,optional"`
}

type Attachment struct {
	URL         string        `json:"url"`
	MIMEType    string        `json:"mime_type"`
	Title       string        `json:"title,optional"`
	SizeInBytes int           `json:"size_in_bytes,optional"`
	Duration    time.Duration `json:"duration_in_seconds,optional"`
}
