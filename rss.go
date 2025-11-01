package main

import (
	"context"
	"net/http"
	"time"
	"io"
	"encoding/xml"
	"html"
	
	"fmt"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err != nil {
			return nil, ctx.Err()
		}
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss []RSSFeed
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		return nil, err
	}

	for i, _ := range rss {
		rss[i].unescapeString()
	}

	fmt.Println(rss)

	return nil, nil
}

// Removes escaped HTML characters from an RSS feed.
func (r *RSSFeed) unescapeString() {
	r.Channel.Title = html.UnescapeString(r.Channel.Title)
	r.Channel.Description = html.UnescapeString(r.Channel.Description)

	for i, _ := range r.Channel.Item {
		r.Channel.Item[i].Title = html.UnescapeString(r.Channel.Item[i].Title)
		r.Channel.Item[i].Description = html.UnescapeString(r.Channel.Item[i].Description)
	}
}