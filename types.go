package jsonfeed

import "time"

type Feed struct {
	Version     string `json:"version" bson:"version"`
	Title       string `json:"title" bson:"title"`
	HomeURL     string `json:"home_page_url" bson:"home_page_url"`
	FeedURL     string `json:"feed_url" bson:"feed_url"`
	Description string `json:"description" bson:"description"`
	UserComment string `json:"user_comment" bson:"user_comment"`
	NextURL     string `json:"next_url" bson:"next_url"`
	Icon        string `json:"icon" bson:"icon"`
	Favicon     string `json:"favicon" bson:"favicon"`
	Expired     bool   `json:"expired" bson:"expired"`
	Items       []Item `json:"items" bson:"items"`
	Hubs        map[string]interface{}
}

type Author struct {
	Name    string `json:"name,optional" bson:"name,optional"`
	SiteURL string `json:"url,optional" bson:"url,optional"`
	Avitar  string `json:"avitar,optional" bson:"avitar,optional"`
}

type Hub struct {
	Type string `json:"type" bson:"type"`
	URL  string `json:"url" bson:"url"`
}

type Item struct {
	ID             string       `json:"id" bson:"_id"`
	URL            string       `json:"url,optional" bson:"url,optional"`
	ExternalURL    string       `json:"external_url,optional" bson:"external_url,optional"`
	Title          string       `json:"title,optional" bson:"title,optional"`
	ContentHTML    string       `json:"content_html,optional" bson:"content_html,optional"`
	ContentText    string       `json:"content_text,optional" bson:"content_text,optional"`
	Summary        string       `json:"summary,optional" bson:"summary,optional"`
	ImageURL       string       `json:"image,optional" bson:"image,optional"`
	BannerImageURL string       `json:"banner_image,optional" bson:"banner_image,optional"`
	DatePublished  time.Time    `json:"date_published,optional" bson:"date_published,optional"`
	DateModified   time.Time    `json:"date_modified,optional" bson:"date_modified,optional"`
	Author         Author       `json:"author,optional" bson:"author,optional"`
	Tags           []string     `json:"tags,optional" bson:"tags,optional"`
	Attachments    []Attachment `json:"attachments,optional" bson:"attachments,optional"`
}

type Attachment struct {
	URL         string        `json:"url" bson:"url"`
	MIMEType    string        `json:"mime_type" bson:"mime_type"`
	Title       string        `json:"title,optional" bson:"title,optional"`
	SizeInBytes int           `json:"size_in_bytes,optional" bson:"size_in_bytes,optional"`
	Duration    time.Duration `json:"duration_in_seconds,optional" bson:"duration_in_seconds,optional"`
}
