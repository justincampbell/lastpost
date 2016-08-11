package main

import (
	"errors"
	"flag"
	"fmt"
	"html"
	"io"
	"os"
	"regexp"
	"time"

	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/justincampbell/timeago"
)

const usage = `lastpost v1.0.0 - github.com/justincampbell/lastpost
usage: lastpost URL`

func main() {
	flag.Parse()
	url := flag.Arg(0)

	if url == "" {
		fmt.Println(usage)
		os.Exit(1)
	}

	feed := rss.New(5, true, nil, itemHandler)
	if err := feed.Fetch(url, charsetReader); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func itemHandler(_ *rss.Feed, _ *rss.Channel, items []*rss.Item) {
	item := items[0]
	c, err := extractContent(item)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c = formatContent(c)

	t, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
	if err == nil {
		fmt.Printf("%s (%s)\n\n%s", item.Title, timeago.FromTime(t), c)
	} else {
		fmt.Printf("%s\n\n%s", item.Title, c)
	}

}

func charsetReader(_ string, r io.Reader) (io.Reader, error) {
	return r, nil
}

func extractContent(item *rss.Item) (string, error) {
	if item == nil {
		return "", errors.New("No item")

	}

	if item.Content != nil && item.Content.Text != "" {
		return item.Content.Text, nil
	}

	if item.Description != "" {
		return item.Description, nil
	}

	return "", errors.New("No content or description found")
}

func formatContent(c string) string {
	c = html.UnescapeString(c)
	c = fmt.Sprintf("\n%s", c)

	for _, pair := range [][]string{
		[]string{"&nbsp;", " "},
		[]string{"&#8230;", "..."},
		[]string{"&#8211;", "-"},
		[]string{"&#215;", "×"},
		[]string{"</\\w+(?:.*)>", "\n"},
		[]string{"<\\w+(?:.*)>", ""},
		[]string{"\n+", "\n"},
		[]string{"\\s\\s+", " "},
		[]string{"^\n*", ""},
		[]string{"\n*$", "\n"},
	} {
		re, replacement := pair[0], pair[1]
		c = regexp.MustCompile(re).ReplaceAllString(c, replacement)
	}

	return c
}
